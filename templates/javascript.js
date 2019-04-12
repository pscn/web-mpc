// short for document.getElementById(id)
var e = function(id) {
  return document.getElementById(id);
};
// short for document.getElementById(id).classList.add(class)
var addcls = function(id, cls) {
  console.log("addcls(" + id + "," + cls + ")");
  e(id).classList.add(cls);
};
// short for document.getElementById(id).classList.rm(class)
var rmcls = function(id, cls) {
  console.log("rmcls(" + id + "," + cls + ")");
  e(id).classList.remove(cls);
};
var hide = function(id) {
  console.log("hide(" + id + ")");
  addcls(id, "hide");
};
var show = function(id) {
  console.log("show(" + id + ")");
  rmcls(id, "hide");
};
var disable = function(id) {
  e(id).disabled = "disabled";
};
var enable = function(id) {
  e(id).disabled = "";
};

/*
 * switch betwwen different views
 */
const views = {
  // view name: button ID
  viewPlaylist: "list",
  viewSearch: "search",
  viewDirectory: "browser"
};
var showView = function(view) {
  return function() {
    for (var k in views) {
      // disable/enable view buttons & show/hide view
      switch (k) {
        case view: // show matching view
          show(k);
          e(views[k]).disabled = "disabled";
          break;
        default:
          // hidde others
          hide(k);
          e(views[k]).disabled = "";
          break;
      }
    }
    switch (
      view // special actions based on the select view
    ) {
      case "viewDirectory":
        command("browse", ""); // send a command
        break;
      case "viewSearch":
        if (view == "viewSearch") {
          e("searchText").select();
          e("searchText").focus();
        }
        break;
    }
  };
};

window.addEventListener("load", function(evt) {
  var ws_addr = e("ws").value; // read from hidden input field
  var ws;

  // send a command on the websocket
  var command = function(cmd, data) {
    if (!ws) {
      return false;
    }
    // console.log('SEND: ' + JSON.stringify(myJson))
    console.log({ cmd, data });
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

  // a few 'globals' to track the process and current state of the player
  var gDuration = 1.0;
  var gElapsed = 0.0;
  var gState = {
    play: "pause",
    consume: false,
    repeat: false,
    single: false,
    random: false,
    song: 0,
    nextsong: 0
  };
  var gErrors = [];
  var gPlaylistFiles = [];
  var readableSeconds = function(value) {
    var min = parseInt(value / 60);
    var sec = parseInt(value % 60);
    if (sec < 10) {
      sec = "0" + sec;
    }
    if (min < 10) {
      min = "&nbsp;&nbsp;" + min;
    } else if (min < 100) {
      min = "&nbsp;" + min;
    }
    return min + ":" + sec;
  };
  // this functions runs forever and gets called every second to update the
  // elapsed and duration information of the active song
  var updateProgress = function() {
    e("elapsed").innerHTML = readableSeconds(gElapsed);
    e("duration").innerHTML = readableSeconds(gDuration);
    if (gState["play"] == "play" && gElapsed < gDuration) {
      // increment the seconds if playing and not finished
      gElapsed += 1;
    }
    if (parseInt(gElapsed) % 2 == 0) {
      e("pause").classList.add("blink");
    } else {
      e("pause").classList.remove("blink");
    }
    setTimeout(updateProgress, 1000);
  };

  var showError = function(msg) {
    if (msg != "") {
      gErrors.push(msg);
    }
    // FIXME: looks ugly
    var sep = "";
    var str = "";
    gErrors.map(function(value) {
      str += sep + value;
      sep = "<br />";
    });
    e("mainError").innerHTML = str;
    show("mainError");

    setTimeout(function() {
      if (gErrors.length > 0) {
        gErrors.shift();
        showError("");
      }
      if (gErrors.length == 0) {
        hideError();
      }
    }, 5000);
  };
  var hideError = function() {
    e("mainError").innerHTML = "";
    hide("mainError");
  };

  var updateStatus = function(data) {
    // FIXME: use nextsong (or nextsongid) to highlight the next song
    const {
      state,
      duration,
      elapsed,
      consume,
      repeat,
      random,
      single,
      song,
      nextsong
    } = data;
    gState = {
      play: state,
      consume: consume,
      repeat: repeat,
      single: single,
      random: random,
      song: song,
      nextsong: nextsong
    };
    // console.log(`updateStatus(${state})`);
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
    // update the mode ctrl
    ["consume", "repeat", "single", "random"].map(function(value) {
      if (gState[value]) {
        show(value + "Disable");
        hide(value + "Enable");
      } else {
        hide(value + "Disable");
        show(value + "Enable");
      }
    });
  };
  var addEvent = function(el, type, fn) {
    if (el.addEventListener) el.addEventListener(type, fn, false);
    else el.attachEvent("on" + type, fn);
  };
  var fitText = function(el, minfs, maxfs) {
    //https://github.com/adactio/FitText.js/blob/master/fittext.js
    var fit = function(el) {
      var resizer = function() {
        el.style.fontSize =
          Math.max(
            Math.min(
              (el.clientWidth / el.innerHTML.length) * 2,
              parseFloat(maxfs)
            ),
            parseFloat(minfs)
          ) + "px";
      };
      // Call once to set.
      resizer();
      // FIXME: add the resize to the element once
      addEvent(window, "resize", resizer);
      addEvent(window, "orientationchange", resizer);
    };

    if (el.length) for (var i = 0; i < el.length; i++) fit(el[i]);
    else fit(el);

    // return set of elements
    return el;
  };

  var updateActiveSong = function(data) {
    const { file, artist, title, album_artist, album } = data;
    e("ctrlSong").title = file;
    e("artist").innerHTML = artist;
    e("title").innerHTML = title;
    //e("title").innerHTML = title;
    if (artist != album_artist) {
      // only show album artist if it's different
      show("albumArtist");
      e("albumArtist").innerHTML = "[" + album_artist + "]&nbsp;";
    } else {
      hide("albumArtist");
    }
    e("album").innerHTML = album;
    // FIXME: do this once and trigger after setting the values
    fitText(e("title"), 16, 48);
    fitText(e("artist"), 16, 38);
    fitText(e("album"), 16, 38);
    fitText(e("albumArtist"), 16, 38);
  };

  var newSongNode = function(id, entry, nr) {
    const { file, artist, title, album, album_artist, duration } = entry;
    var node = e(id).cloneNode(true);
    node.classList.remove("hide");
    node.title = file;
    node.id = id + nr;
    node.querySelector("#songCellArtist").innerHTML = artist;
    node.querySelector("#songCellTitle").innerHTML = title;
    node.querySelector("#songCellAlbum").innerHTML = album;
    if (artist != album_artist) {
      // only show album artist if it's different
      node.querySelector("#songCellAlbumArtist").innerHTML =
        "[" + album_artist + "]&nbsp;";
    } else {
      node.querySelector("#songCellAlbumArtist").classList.add("hide");
    }
    node.querySelector("#songCellArtist").innerHTML = artist;
    node.querySelector("#songCellDuration").innerHTML = readableSeconds(
      duration
    );
    fitText(node.querySelector("#songCellArtist"), 16, 28);
    fitText(node.querySelector("#songCellTitle"), 16, 28);
    fitText(node.querySelector("#songCellAlbum"), 16, 28);
    fitText(node.querySelector("#songCellAlbumArtist"), 16, 28);
    return node;
  };

  var processResponse = function(obj) {
    const { type, data } = obj;
    console.log({ type, data });

    switch (type) {
      case ("error", "info"):
        showError(data);
        break;
      case "update":
        updateStatus(data.status);

        e("playlist").innerHTML = ""; // delete old playlist
        gPlaylistFiles.length = 0;
        // used to figure out if we need to show prio buttons or not:
        // if state == stop && activeFile is set && queuesize <=2 → no need to show buttons
        // if state == stop && activeFile is not set → show buttons
        // if state != stop && queuesize <=2 → do not show buttons
        // FIXME: is that correct?
        const queueSize = data.queue.length;
        const state = data.status.state;
        const activeFile = data.activeSong.file;
        data.queue.map(function(entry, i) {
          const { file, prio, position, isActive, isNext } = entry;
          if (isActive) {
            return; // don't show the active song in the list
          }
          gPlaylistFiles.push(file);
          var node = newSongNode("playlistEntry", entry, i);
          const btnPlay = node.querySelector("#plPlay");
          if (isActive) {
            btnPlay.disabled = "disabled";
            node.classList.add("activeSong");
          } else {
            btnPlay.onclick = btnCommand("play", position.toString());
          }
          if (isNext) {
            node.classList.add("nextSong");
          }
          node.querySelector("#plRemove").onclick = btnCommand(
            "remove",
            position.toString()
          );
          // FIXME: this only makes sense in certain modes (random + ?)
          const btnPrio1 = node.querySelector("#plPrio1");
          const btnPrio2 = node.querySelector("#plPrio2");
          const btnPrio3 = node.querySelector("#plPrio3");
          if (!gState["random"]) {
            btnPrio1.classList.add("hide");
            btnPrio2.classList.add("hide");
            btnPrio3.classList.add("hide");
          } else {
            // see explanation above
            if (
              (state != "stop" || activeFile != "") &&
              (isActive || queueSize <= 2)
            ) {
              btnPrio1.disabled = "disabled";
              btnPrio2.disabled = "disabled";
              btnPrio3.disabled = "disabled";
            } else {
              switch (prio) {
                case 0:
                  btnPrio3.disabled = "disabled";
                  break;
                case 127:
                  btnPrio2.disabled = "disabled";
                  break;
                case 255:
                  btnPrio1.disabled = "disabled";
                  break;
              }
              btnPrio1.onclick = btnCommand(
                "prio",
                "255:" + position.toString()
              );
              btnPrio2.onclick = btnCommand(
                "prio",
                "127:" + position.toString()
              );
              btnPrio3.onclick = btnCommand("prio", "0:" + position.toString());
            }
          }
          e("playlist").append(node);
        });
        window.dispatchEvent(new Event("resize")); // trigger resize events on the song stuff

        updateActiveSong(data.activeSong);
        break;
      case "searchResult":
        e("searchResult").innerHTML = ""; // delete old search result
        // console.log({ gFiles: gPlaylistFiles });
        if (data.truncated) {
          showError("search result limited to " + data.maxResults);
        }
        if (data.searchResult.length == 0) {
          showError("nothing found");
        }
        data.searchResult.map(function(entry) {
          const { file } = entry;
          const node = newSongNode("searchEntry", entry);
          const btn = node.querySelector("#srAdd");
          // disable button for files already in the playlist
          if (gPlaylistFiles.includes(file)) {
            // FIXME: should we add a button to remove it from the playlist?
            btn.disabled = "disabled";
          }
          btn.onclick = function() {
            btn.disabled = "disabled";
            return command("add", file);
          };
          e("searchResult").append(node);
        });

        window.dispatchEvent(new Event("resize")); // trigger resize events on the song stuff

        break;
      case "directoryList":
        e("directoryList").innerHTML = "";
        node = e("directoryListEntry").cloneNode(true);
        node.id = "dlRowParent";
        node.title = data.parent;
        node.classList.remove("hide");
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
            node.classList.remove("hide");
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
            node.classList.remove("hide");
            node.querySelector("#srArtist").innerHTML = entry.artist;
            node.querySelector("#srTitle").innerHTML = entry.title;
            node.querySelector("#srAlbum").innerHTML = entry.album;
            if (entry.artist != entry.album_artist) {
              node.querySelector("#srAlbumArtist").innerHTML =
                "[" + entry.album_artist + "]&nbsp;";
            } else {
              node.querySelector("#srAlbumArtist").classList.add("hide");
            }
            node.querySelector("#srArtist").innerHTML = entry.artist;
            node.querySelector("#srDuration").innerHTML = readableSeconds(
              entry.duration
            );
            {
              const file = entry.file;
              const node = node.querySelector("#srAdd");
              node.onclick = function(evt) {
                node.disabled = "disabled";
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
    command("updateRequest", "");
  };
  updateProgress();

  var stop = function() {
    console.log("stop");
    ["play"].map(show);
    ["pause", "resume"].map(hide);
    ["stop", "next", "previous"].map(disable);
  };
  var play = function() {
    console.log("play");
    ["pause"].map(show);
    ["resume", "play"].map(hide);
    ["stop", "next", "previous"].map(enable);
  };
  var pause = function() {
    console.log("pause");
    ["resume"].map(show);
    ["pause", "play"].map(hide);
    ["stop", "next", "previous"].map(enable);
  };
  var togglePlayPause = function(state) {
    // console.log(`togglePlayPause(${state})`);
    switch (state) {
      case "play":
        play();
        break;
      case "pause":
        pause();
        break;
    }
  };

  // add onclick to every button
  for (var k in views) {
    e(views[k]).onclick = showView(k);
  }
  ["random", "consume", "repeat", "single"].map(function(value) {
    e(value + "Enable").onclick = btnCommand(value, "enable");
    e(value + "Disable").onclick = btnCommand(value, "disable");
  });

  /*
   * onclick function assignments
   */
  e("submitSearch").onclick = function(evt) {
    return command("search", e("searchText").value);
  };

  e("searchText").onchange = function(evt) {
    return command("search", e("searchText").value);
  };
  e("connect").onclick = openWebSocket;

  // add onclick function for all controls
  ["play", "resume", "pause", "stop", "next", "previous"].map(function(value) {
    e(value).onclick = function(evt) {
      console.log(`Control: ${value}`);
      return command(value, "");
    };
  });

  // show the viewPlaylist
  show("viewPlaylist")();
});

// eof
