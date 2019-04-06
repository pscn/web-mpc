window.addEventListener("load", function(evt) {
  // short for document.getElementById
  var e = function(id) {
    return document.getElementById(id);
  };
  // helpers to easily show/hide or enable/disable elements
  // example: ["play"].map(eShow) // set style.display for "play" to ""
  var eHide = function(id) {
    e(id).style.display = "none";
  };
  var eShow = function(id) {
    e(id).style.display = "";
  };
  var eDisable = function(id) {
    e(id).disabled = "disabled";
  };
  var eEnable = function(id) {
    e(id).disabled = "";
  };

  var ws_addr = e("ws").value; // read from hidden input field
  var ws;

  // send a command on the websocket
  var command = function(cmd, data) {
    if (!ws) {
      return false;
    }
    // console.log('SEND: ' + JSON.stringify(myJson))
    ws.send(JSON.stringify({ command: cmd, data: data }));
    return true;
  };

  // wrapper to use command as an onclick
  // example: e("play").onclick = btnCommand("play", 1)
  var btnCommand = function(cmd, data) {
    return function() {
      return command(cmd, data);
    };
  };

  // https://www.w3schools.com/js/js_cookies.asp
  // gets the value of the cookie cookieName
  var getCookie = function(cookieName) {
    var name = cookieName + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(";");
    for (var i = 0; i < ca.length; i++) {
      var c = ca[i];
      while (c.charAt(0) == " ") {
        c = c.substring(1);
      }
      if (c.indexOf(name) == 0) {
        return c.substring(name.length, c.length);
      }
    }
    return "";
  };

  var updateConfig = function() {
    const mpdAddr = getCookie("mpd");
    if (mpdAddr != "") {
      const parts = mpdAddr.split(":");
      // having fun with javascript, this does:
      // e("configHost").value = parts[0]
      // e("configPort").value = parts[1] etc
      ["configHost", "configPort", "configPass"].map(
        (id, i) => (e(id).value = parts[i])
      );
    }
  };
  updateConfig();

  // a few 'globals' to track the process and current state of the player
  var gDuration = 1.0;
  var gElapsed = 0.0;
  var gState = "pause";
  var readableSeconds = function(value) {
    const min = parseInt(value / 60);
    var sec = parseInt(value % 60);
    if (sec < 10) {
      sec = "0" + sec;
    }
    return min + ":" + sec;
  };
  // this functions runs forever and gets called every second to update the
  // elapsed and duration information of the active song
  var updateProgress = function() {
    e("elapsed").innerHTML = readableSeconds(gElapsed);
    e("duration").innerHTML = readableSeconds(gDuration);
    if (gState == "play" && gElapsed < gDuration) {
      // increment the seconds if playing and not finished
      gElapsed += 1.0;
    }
    setTimeout(updateProgress, 1000);
  };

  var showError = function(msg) {
    // FIXME: looks ugly
    e("error").innerHTML =
      e("error").innerHTML == "" ? msg : e("error").innerHTML + "<br />" + msg;
    eShow("error");
  };
  var hideError = function() {
    e("error").innerHTML = "";
    eHide("error");
  };

  var updateStatus = function(data) {
    const { state, duration, elapsed, consume, repeat, random, single } = data;
    gState = state;
    console.log(`updateStatus(${state})`);
    switch (state) {
      case "pause":
      case "play":
        togglePlayPause(state);
        gDuration = duration;
        gElapsed = elapsed;
        break;
      case "stop":
        stop();
        gDuration = 0.0;
        gElapsed = 0.0;
        break;
    }
    // update the config view
    e("consume").checked = consume ? "checked" : "";
    e("repeat").checked = repeat ? "checked" : "";
    e("random").checked = random ? "checked" : "";
    e("single").checked = single ? "checked" : "";
  };
  var updateActiveSong = function(data) {
    const { file, artist, title, album_artist, album } = data;
    e("activeSong").title = file;
    e("artist").innerHTML = artist;
    e("title").innerHTML = title;
    if (artist != album_artist) {
      // only show album artist if it's different
      eShow("albumArtist");
      e("albumArtist").innerHTML = "[" + album_artist + "]&nbsp;";
    } else {
      eHide("albumArtist");
    }
    e("album").innerHTML = album;
  };

  var newSongNode = function(id, entry) {
    const { file, artist, title, album, album_artist, duration } = entry;
    var node = e(id).cloneNode(true);
    node.style.display = "";
    node.title = file;
    node.querySelector("#songCellArtist").innerHTML = artist;
    node.querySelector("#songCellTitle").innerHTML = title;
    node.querySelector("#songCellAlbum").innerHTML = album;
    if (artist != album_artist) {
      // only show album artist if it's different
      node.querySelector("#songCellAlbumArtist").innerHTML =
        "[" + album_artist + "]&nbsp;";
    } else {
      node.querySelector("#songCellAlbumArtist").style.display = "none";
    }
    node.querySelector("#songCellArtist").innerHTML = artist;
    node.querySelector("#songCellDuration").innerHTML = readableSeconds(
      duration
    );
    return node;
  };

  var processResponse = function(obj) {
    console.log({ obj });
    const { type, data } = obj;

    switch (type) {
      case ("error", "info"):
        showError(data);
        break;
      case "status":
        updateStatus(data);
        break;
      case "activeSong":
        updateActiveSong(data);
        break;
      case "activePlaylist":
        e("playlist").innerHTML = ""; // delete old playlist
        data.Playlist.map(function(entry, i) {
          const { file } = entry;
          var node = newSongNode("playlistEntry", entry);
          // disable the play button for the active song
          node.querySelector("#plPlay").disabled =
            file == e("activeSong").title ? "disabled" : "";
          node.querySelector("#plPlay").onclick = btnCommand(
            "play",
            i.toString()
          );
          node.querySelector("#plRemove").onclick = btnCommand(
            "remove",
            i.toString()
          );
          e("playlist").append(node);
        });
        break;
      case "searchResult":
        e("searchResult").innerHTML = ""; // delete old search result
        data.SearchResult.map(function(entry) {
          const { file } = entry;
          var node = newSongNode("searchEntry", entry);
          node.querySelector("#srAdd").onclick = btnCommand("add", file);
          e("searchResult").append(node);
        });
        break;
      case "directoryList":
        e("directoryList").innerHTML = "";
        node = e("directoryListEntry").cloneNode(true);
        node.id = "dlRowParent";
        node.title = data.parent;
        node.style.display = "";
        node.querySelector("#dlName").innerHTML = data.parent;
        node.querySelector("#dlBrowse").onclick = btnCommand(
          "browse",
          data.parent
        );
        e("directoryList").append(node);
        data.directoryList.map(function(entry, i) {
          var node;
          if (entry.type == "directory") {
            node = e("directoryListEntry").cloneNode(true);
            node.id = "dlRow" + i;
            node.style.display = "";
            node.querySelector("#dlName").innerHTML = entry.directory;

            {
              const name = entry.directory;
              node.querySelector("#dlBrowse").onclick = function(evt) {
                return command("browse", name);
              };
            }
          } else {
            node = e("searchEntry").cloneNode(true);
            node.id = "srRow" + i;
            node.style.display = "";
            node.querySelector("#srArtist").innerHTML = entry.artist;
            node.querySelector("#srTitle").innerHTML = entry.title;
            node.querySelector("#srAlbum").innerHTML = entry.album;
            if (entry.artist != entry.album_artist) {
              node.querySelector("#srAlbumArtist").innerHTML =
                "[" + entry.album_artist + "]&nbsp;";
            } else {
              node.querySelector("#srAlbumArtist").style.display = "none";
            }
            node.querySelector("#srArtist").innerHTML = entry.artist;
            node.querySelector("#srDuration").innerHTML = readableSeconds(
              entry.duration
            );
            {
              const file = entry.file;
              node.querySelector("#srAdd").onclick = function(evt) {
                return command("add", file);
              };
            }
          }
          e("directoryList").append(node);
        });
        break;
    }
  };
  var openWebSocket = function() {
    ws = new WebSocket(ws_addr);
    ws.onopen = function(evt) {
      console.log("OPEN");
      hideError();
      e("connect").disabled = "disabled";
    };
    ws.onclose = function(evt) {
      console.log("CLOSE");
      ws = null;
      showError("no connection");
      e("connect").disabled = "";
    };
    ws.onmessage = function(evt) {
      processResponse(JSON.parse(evt.data));
    };
    ws.onerror = function(evt) {
      showError(evt.data);
    };
  };
  openWebSocket();

  window.onfocus = function(event) {
    // request a fresh status as some browsers (e. g. Chrome) suspend our
    // progress setTimeout functions
    command("statusRequest", "");
  };
  updateProgress();

  var stop = function() {
    ["play"].map(eShow);
    ["pause", "resume"].map(eHide);
    ["stop", "next", "previous"].map(eDisable);
  };
  var play = function() {
    ["pause"].map(eShow);
    ["resume", "play"].map(eHide);
    ["stop", "next", "previous"].map(eEnable);
  };
  var pause = function() {
    ["resume"].map(eShow);
    ["pause", "play"].map(eHide);
    ["stop", "next", "previous"].map(eEnable);
  };
  var togglePlayPause = function(state) {
    console.log(`togglePlayPause(${state})`);
    switch (state) {
      case "play":
        play();
        break;
      case "pause":
        pause();
        break;
    }
  };

  /*
   * switch betwwen different views
   */
  var views = {
    // view name: button ID
    playlistView: "list",
    searchView: "search",
    folderView: "browser",
    configView: "config"
  };
  var show = function(view) {
    return function() {
      switch (view) {
        case "folderView":
          command("browse", "");
          break;
      }
      for (var k in views) {
        switch (k) {
          case view: // show matching view
            e(k).style.display = "";
            e(views[k]).disabled = "disabled";
            break;
          default:
            // hidde others
            e(k).style.display = "none";
            e(views[k]).disabled = "";
            break;
        }
      }
    };
  };
  // add onclick to every button
  for (var k in views) {
    e(views[k]).onclick = show(k);
  }

  /*
   * onclick function assignments
   */
  e("submitSearch").onclick = function(evt) {
    return command("search", e("searchText").value);
  };

  e("submitConfig").onclick = function(evt) {
    // update the cookie and reconnect to the websocket
    // which will read the cookie and apply the changes
    document.cookie =
      "mpd=" +
      e("configHost").value +
      ":" +
      e("configPort").value +
      ":" +
      e("configPass").value;
    if (ws != null) {
      ws.close();
    }
    setTimeout(openWebSocket, 1000);
  };
  e("searchText").onchange = function(evt) {
    return command("search", e("searchText").value);
  };
  e("connectionStatus").onclick = openWebSocket;

  // add onclick function for all controls
  ["play", "resume", "pause", "stop", "next", "previous"].map(function(value) {
    e(value).onclick = function(evt) {
      console.log(`Control: ${value}`);
      return command(value, "");
    };
  });

  // show the playlistView
  show("playlistView");
});

// eof
