window.addEventListener("load", function (evt) {
  var ws;

  var btnPlay = document.getElementById("play");
  var btnPause = document.getElementById("pause");
  var btnResume = document.getElementById("resume");
  var btnStop = document.getElementById("stop");
  var btnNext = document.getElementById("next");
  var btnPrevious = document.getElementById("previous");
  var pause = function () {
    btnPlay.style.display = "none";
    btnPause.style.display = "none";
    btnResume.style.display = "";
    btnStop.style.display = ""; btnStop.disabled = "";
    btnNext.style.display = ""; btnNext.disabled = "";
    btnPrevious.style.display = ""; btnPrevious.disabled = "";
  };

  var play = function () {
    btnPlay.style.display = "none";
    btnPause.style.display = "";
    btnResume.style.display = "none";
    btnStop.style.display = ""; btnStop.disabled = "";
    btnNext.style.display = ""; btnNext.disabled = "";
    btnPrevious.style.display = ""; btnPrevious.disabled = "";
  };

  var stop = function () {
    btnPlay.style.display = "";
    btnPause.style.display = "none";
    btnResume.style.display = "none";
    btnStop.style.display = ""; btnStop.disabled = "disabled";
    btnNext.style.display = ""; btnNext.disabled = "disabled";
    btnPrevious.style.display = ""; btnPrevious.disabled = "disabled";
  };

  var cs = document.getElementById("cs");
  var csTitle = document.getElementById("csTitle");
  var csArtist = document.getElementById("csArtist");
  var csAlbumArtist = document.getElementById("csAlbumArtist");
  var csAlbum = document.getElementById("csAlbum");
  var csElapsed = document.getElementById("csElapsed");
  var csDuration = document.getElementById("csDuration");

  var currentPlaylist = document.getElementById("playlist");
  var searchResult = document.getElementById("searchResult");

  var duration = 1.0;
  var elapsed = 0.0;
  var state = "pause";
  var readableSeconds = function (value) {
    var min = parseInt(value / 60);
    var sec = parseInt(value % 60);
    if (sec < 10) { sec = "0" + sec; }
    return min + ":" + sec;
  };
  var updateProgress = function () {
    // console.log("progress: " + duration + " / " + elapsed);
    csElapsed.innerHTML = readableSeconds(elapsed);
    csDuration.innerHTML = readableSeconds(duration);
    if ((state == "play") && (elapsed < duration)) { elapsed += 1.0; }
    setTimeout(updateProgress, 1000);
  };

  ws_addr = document.getElementById("ws").value;
  ws = new WebSocket(ws_addr);
  ws.onopen = function (evt) { console.log("OPEN"); }
  ws.onclose = function (evt) { console.log("CLOSE"); ws = null; }
  ws.onmessage = function (evt) {
    console.log("RESPONSE: " + evt.data);
    obj = JSON.parse(evt.data);
    if (obj.type == 1) {
    } else if (obj.type == 2) {
      if (obj.data.state == "pause") {
        pause();
        state = "pause";
        duration = obj.data.duration;
        elapsed = obj.data.elapsed;
      } else if (obj.data.state == "play") {
        play();
        state = "play";
        duration = obj.data.duration;
        elapsed = obj.data.elapsed;
      } else if (obj.data.state == "stop") {
        stop();
        state = "stop";
        duration = 0.0;
        elapsed = 0.0;
      }
    } else if (obj.type == 3) {
      cs.title = obj.data.file;
      csArtist.innerHTML = obj.data.artist;
      csTitle.innerHTML = obj.data.title;
      if (obj.data.artist != obj.data.album_artist) {
        csAlbumArtist.style.display = "";
        csAlbumArtist.innerHTML = "[" + obj.data.album_artist + "]&nbsp;";
      } else {
        csAlbumArtist.style.display = "none";
      }
      csAlbum.innerHTML = obj.data.album;
    } else if (obj.type == 4) {
      console.log("currentPlaylist")
      currentPlaylist.innerHTML = "";
      var list = "";
      for (var i = 0; i < obj.data.Playlist.length; i++) {
        var playlistEntry = document.getElementById("playlistEntry")
        var node = playlistEntry.cloneNode(true);
        node.id = "plRow" + i;
        node.style.display = "";
        node.querySelector("#plArtist").innerHTML = obj.data.Playlist[i].artist;
        node.querySelector("#plTitle").innerHTML = obj.data.Playlist[i].title;
        node.querySelector("#plAlbum").innerHTML = obj.data.Playlist[i].album;
        if (obj.data.Playlist[i].artist != obj.data.Playlist[i].album_artist) {
          node.querySelector("#plAlbumArtist").innerHTML = "[" + obj.data.Playlist[i].album_artist + "]&nbsp;";
        } else {
          node.querySelector("#plAlbumArtist").style.display = "none";
        }
        node.querySelector("#plArtist").innerHTML = obj.data.Playlist[i].artist;
        node.querySelector("#plArtist").innerHTML = obj.data.Playlist[i].artist;
        {
          const j = i;
          const file = obj.data.Playlist[i].file;
          if (file == cs.title) {
            node.querySelector("#plPlay").disabled = "disabled";
          }
          node.querySelector("#plPlay").onclick = function (evt) {
            return command("play" + j);
          };
          node.querySelector("#plRemove").onclick = function (evt) {
            return command("remove" + j);
          };
        }
        currentPlaylist.append(node);
      }
    } else if (obj.type == 5) {
      console.log("searchResult")
      searchResult.innerHTML = "";
      var list = "";
      for (var i = 0; i < obj.data.Playlist.length; i++) {
        var playlistEntry = document.getElementById("searchEntry")
        var node = playlistEntry.cloneNode(true);
        node.id = "srRow" + i;
        node.style.display = "";
        node.querySelector("#srArtist").innerHTML = obj.data.Playlist[i].artist;
        node.querySelector("#srTitle").innerHTML = obj.data.Playlist[i].title;
        node.querySelector("#srAlbum").innerHTML = obj.data.Playlist[i].album;
        if (obj.data.Playlist[i].artist != obj.data.Playlist[i].album_artist) {
          node.querySelector("#srAlbumArtist").innerHTML = "[" + obj.data.Playlist[i].album_artist + "]&nbsp;";
        } else {
          node.querySelector("#srAlbumArtist").style.display = "none";
        }
        node.querySelector("#srArtist").innerHTML = obj.data.Playlist[i].artist;
        node.querySelector("#srArtist").innerHTML = obj.data.Playlist[i].artist;
        {
          const file = obj.data.Playlist[i].file;
          node.querySelector("#srAdd").onclick = function (evt) {
            return command("add" + file);
          };
        }
        searchResult.append(node);
      }
    }
  }
  ws.onerror = function (evt) {
    console.log("ERROR: " + evt.data);
  }
  window.onfocus = function (event) {
    // request a fresh status as some browsers (e. g. Chrome) suspend our
    // progress bar setTimeout functions
    command("status");
  }
  updateProgress();

  var command = function (cmd) {
    if (!ws) { return false; }
    myJson = { "command": cmd }
    console.log("SEND: " + myJson);
    ws.send(JSON.stringify(myJson));
    return false;
  }

  // add onclick function for all controls
  var controls = ["play", "resume", "pause", "stop", "next", "previous"]
  controls.forEach(activator);
  function activator(value) {
    document.getElementById(value).onclick = function (evt) {
      return command(value);
    };
  }
  var showList = function (evt) {
    document.getElementById("playlist").style.display = "";
    document.getElementById("searchBox").style.display = "none";
    document.getElementById("searchResult").style.display = "none";
    document.getElementById("search").disabled = "";
    document.getElementById("list").disabled = "disabled";
  }
  var showSearch = function (evt) {
    document.getElementById("playlist").style.display = "none";
    document.getElementById("searchBox").style.display = "";
    document.getElementById("searchText").focus();
    document.getElementById("searchResult").innerHTML = "";
    document.getElementById("searchResult").style.display = "";
    document.getElementById("search").disabled = "disabled";
    document.getElementById("list").disabled = "";
  }
  document.getElementById("search").onclick = showSearch;

  document.getElementById("list").onclick = showList;
  document.getElementById("closeSearch").onclick = showList;

  document.getElementById("submitSearch").onclick = function (evt) {
    return command("search" + document.getElementById("searchText").value);
  }
  document.getElementById("searchText").onchange = function (evt) {
    return command("search" + document.getElementById("searchText").value);
  }

});