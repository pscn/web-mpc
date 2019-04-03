window.addEventListener("load", function (evt) {
  var ws;

  // hard coded command types: synchronize with command.go
  var cmd = {
    "play": 0,
    "resume": 1,
    "pause": 2,
    "stop": 3,
    "next": 4,
    "previous": 5,
    "add": 6,
    "remove": 7,
    "search": 8,
    "statusrRequest": 9,
  };

  // hard coded event types: synchronize with message.go
  var ev = {
    "error": 0,
    "info": 1,
    "status": 2,
    "currentSong": 3,
    "playlist": 4,
    "searchResult":5,
  };

  // pre-load some document.getElementById calls to have the code a little
  // shorter down the road
  var elementIDs = [
    "error",
    "playlist", "searchBox", "searchText", "searchResult",
    "search", "submitSearch", "closeSearch", "searchText", "list",
    "ws",
    "playlistEntry", "playlist", "searchEntry",
    "cs", "csTitle", "csArtist", "csAlbumArtist", "csAlbum",
    "csElapsed", "csDuration",
    "play", "resume", "pause", "stop", "next", "previous"];
  var el = {}
  elementIDs.map(function (value) {
    el[value] = document.getElementById(value);
  });

  // a few "globals" to track the process and current state of the player
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
    el["csElapsed"].innerHTML = readableSeconds(elapsed);
    el["csDuration"].innerHTML = readableSeconds(duration);
    if ((state == "play") && (elapsed < duration)) {
      elapsed += 1.0;
    }
    setTimeout(updateProgress, 1000);
  };

  ws_addr = el["ws"].value;
  ws = new WebSocket(ws_addr);
  ws.onopen = function (evt) {
    console.log("OPEN");
  }
  ws.onclose = function (evt) {
    console.log("CLOSE"); ws = null;
  }
  ws.onmessage = function (evt) {
    console.log("RESPONSE: " + evt.data);
    obj = JSON.parse(evt.data);
    switch (obj.type) {
      case ev["error"], ev["info"]:
        el["error"].innerHTML = obj.data;
        el["error"].style.display = "";
        console.log(obj.data);
        break;
      case ev["status"]:
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
        break;
      case ev["currentSong"]:
        el["cs"].title = obj.data.file;
        el["csArtist"].innerHTML = obj.data.artist;
        el["csTitle"].innerHTML = obj.data.title;
        if (obj.data.artist != obj.data.album_artist) {
          el["csAlbumArtist"].style.display = "";
          el["csAlbumArtist"].innerHTML = "[" + obj.data.album_artist + "]&nbsp;";
        } else {
          el["csAlbumArtist"].style.display = "none";
        }
        el["csAlbum"].innerHTML = obj.data.album;
        break;
      case ev["playlist"]:
        console.log("playlist")
        el["playlist"].innerHTML = "";
        obj.data.Playlist.map(function (entry, i) {
          var node = el["playlistEntry"].cloneNode(true);
          node.id = "plRow" + i;
          node.style.display = "";
          node.querySelector("#plArtist").innerHTML = entry.artist;
          node.querySelector("#plTitle").innerHTML = entry.title;
          node.querySelector("#plAlbum").innerHTML = entry.album;
          if (entry.artist != entry.album_artist) {
            node.querySelector("#plAlbumArtist").innerHTML = "[" + entry.album_artist + "]&nbsp;";
          } else {
            node.querySelector("#plAlbumArtist").style.display = "none";
          }
          node.querySelector("#plArtist").innerHTML = entry.artist;
          node.querySelector("#plDuration").innerHTML =
            readableSeconds(entry.duration);
          {
            const j = i;
            const file = entry.file;
            if (file == cs.title) {
              node.querySelector("#plPlay").disabled = "disabled";
            }
            node.querySelector("#plPlay").onclick = function (evt) {
              return command("play", j.toString());
            };
            node.querySelector("#plRemove").onclick = function (evt) {
              return command("remove", j.toString());
            };
          }
          el["playlist"].append(node);
        })
        break;
      case ev["searchResult"]:
        el["searchResult"].innerHTML = "";
        obj.data.Playlist.map(function (entry, i) {
          var node = el["searchEntry"].cloneNode(true);
          node.id = "srRow" + i;
          node.style.display = "";
          node.querySelector("#srArtist").innerHTML = entry.artist;
          node.querySelector("#srTitle").innerHTML = entry.title;
          node.querySelector("#srAlbum").innerHTML = entry.album;
          if (entry.artist != entry.album_artist) {
            node.querySelector("#srAlbumArtist").innerHTML = "[" + entry.album_artist + "]&nbsp;";
          } else {
            node.querySelector("#srAlbumArtist").style.display = "none";
          }
          node.querySelector("#srArtist").innerHTML = entry.artist;
          node.querySelector("#srDuration").innerHTML = readableSeconds(entry.duration);
          {
            const file = entry.file;
            node.querySelector("#srAdd").onclick = function (evt) {
              return command("add", file);
            };
          }
          el["searchResult"].append(node);
        })
        break;
    }
  }
  ws.onerror = function (evt) {
    console.log("ERROR: " + evt.data);
  }
  window.onfocus = function (event) {
    // request a fresh status as some browsers (e. g. Chrome) suspend our
    // progress bar setTimeout functions
    command("statusRequest", "");
  }
  updateProgress();

  var command = function (cmdType, data) {
    if (!ws) { return false; }
    myJson = { "command": cmd[cmdType], "data": data };
    console.log("SEND: " + JSON.stringify(myJson));
    ws.send(JSON.stringify(myJson));
    return false;
  }

  // add onclick function for all controls
  var pause = function () {
    el["play"].style.display = "none";
    el["pause"].style.display = "none";
    el["resume"].style.display = "";
    el["stop"].style.display = ""; el["stop"].disabled = "";
    el["next"].style.display = ""; el["next"].disabled = "";
    el["previous"].style.display = ""; el["previous"].disabled = "";
  };

  var play = function () {
    el["play"].style.display = "none";
    el["pause"].style.display = "";
    el["resume"].style.display = "none";
    el["stop"].style.display = ""; el["stop"].disabled = "";
    el["next"].style.display = ""; el["next"].disabled = "";
    el["previous"].style.display = ""; el["previous"].disabled = "";
  };

  var stop = function () {
    el["play"].style.display = "";
    el["pause"].style.display = "none";
    el["resume"].style.display = "none";
    el["stop"].style.display = ""; el["stop"].disabled = "disabled";
    el["next"].style.display = ""; el["next"].disabled = "disabled";
    el["previous"].style.display = ""; el["previous"].disabled = "disabled";
  };

  var buttonIDs = ["play", "resume", "pause", "stop", "next", "previous"];
  buttonIDs.map(function (value) {
    document.getElementById(value).onclick = function (evt) {
      return command(value, "");
    };
  });

  var showList = function (evt) {
    el["playlist"].style.display = "";
    el["searchBox"].style.display = "none";
    el["searchResult"].style.display = "none";
    el["search"].disabled = "";
    el["list"].disabled = "disabled";
  }
  var showSearch = function (evt) {
    el["playlist"].style.display = "none";
    el["searchBox"].style.display = "";
    el["searchText"].focus();
    el["searchText"].select();
    el["searchResult"].innerHTML = "";
    el["searchResult"].style.display = "";
    el["search"].disabled = "disabled";
    el["list"].disabled = "";
  }
  el["search"].onclick = showSearch;

  el["list"].onclick = showList;
  el["closeSearch"].onclick = showList;

  el["submitSearch"].onclick = function (evt) {
    return command("search", el["searchText"].value);
  }
  el["searchText"].onchange = function (evt) {
    return command("search", el["searchText"].value);
  }

  showList();

});