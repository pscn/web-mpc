var log = function(msg) {
  console.log(msg);
};
var debug = function(msg) {
  if (false) log(msg);
};
// short for document.getElementById(id)
var e = function(id) {
  return document.getElementById(id);
};

// https://stackoverflow.com/questions/7444451/how-to-get-the-actual-rendered-font-when-its-not-defined-in-css
var css = function(element, property) {
  return window.getComputedStyle(element, null).getPropertyValue(property);
};

// https://stackoverflow.com/questions/118241/calculate-text-width-with-javascript/21015393#21015393
var getTextWidth = function(text, fontSize) {
  // if given, use cached canvas for better performance
  // else, create new canvas
  var canvas =
    getTextWidth.canvas ||
    (getTextWidth.canvas = document.createElement("canvas"));
  var fontFamily =
    getTextWidth.fontFamily ||
    (getTextWidth.fontFamily = css(e("body"), "font-family"));
  var context = canvas.getContext("2d");
  context.font = fontSize + "px " + fontFamily;
  var metrics = context.measureText(text);
  return metrics.width;
};
// short for document.getElementById(id).classList.add(class)
var addcls = function(id, cls) {
  debug("addcls(" + id + "," + cls + ")");
  e(id).classList.add(cls);
};
// short for document.getElementById(id).classList.rm(class)
var rmcls = function(id, cls) {
  debug("rmcls(" + id + "," + cls + ")");
  e(id).classList.remove(cls);
};
var hide = function(id) {
  debug("hide(" + id + ")");
  addcls(id, "hide");
};
var show = function(id) {
  debug("show(" + id + ")");
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

var addEvent = function(el, type, fn) {
  if (el.addEventListener) {
    el.addEventListener(type, fn, false);
  } else {
    el.attachEvent("on" + type, fn);
  }
};

// Resize voodoo
var resize = function(el, minFS, maxFS) {
  var fs = maxFS;
  var txt = el.innerHTML;
  var maxWidth = el.clientWidth;
  while (fs >= minFS && getTextWidth(txt, fs) > maxWidth) {
    fs -= 2;
  }
  if (fs < minFS) {
    fs = minFS;
    maxWidth *= 2; // we allow one page break
    var truncated = false;
    while (txt.length > 8 && getTextWidth(txt, fs) > maxWidth) {
      log("too long: " + txt);
      txt = txt.slice(0, txt.length - 1);
      truncated = true;
    }
    if (truncated) {
      el.innerHTML = txt.slice(0, txt.length - 3) + "&hellip;";
    }
  }
  el.style.fontSize = fs + "px";
};
var resizer = function() {
  var el = document.getElementsByClassName("resize");
  var i;
  for (i = 0; i < el.length; i++) {
    resize(el[i], 12, 28);
  }
};

var newResize = function(el) {
  var maxFS = el.clientHeight;
};

var triggerResize = function() {
  window.dispatchEvent(new Event("resize")); // trigger resize events on the song stuff
};

// send a command on the websocket
var command = function(cmd, data) {
  if (!ws) return false;
  log({ cmd, data });
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

var ws_addr, ws;

window.addEventListener("load", function(evt) {
  ws_addr = e("ws").value; // read from hidden input field
  addEvent(window, "resize", resizer);
  addEvent(window, "orientationchange", resizer);

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
    /*
    if (min < 10) {
      min = "&nbsp;&nbsp;" + min;
    } else if (min < 100) {
      min = "&nbsp;" + min;
    }
    */
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
    // debug(`updateStatus(${state})`);
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

  var updateActiveSong = function(data) {
    const { file, artist, title, album_artist, album } = data;
    e("ctrlSong").title = file;
    e("artist").innerHTML = artist;
    e("title").innerHTML = title;
    //e("title").innerHTML = title;
    if (artist != album_artist) {
      e("album").innerHTML = album + " [" + album_artist + "]";
    } else {
      e("album").innerHTML = album;
    }
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
      node.querySelector("#songCellAlbum").innerHTML =
        node.querySelector("#songCellAlbum").innerHTML +
        " [" +
        album_artist +
        "]";
    }
    node.querySelector("#songCellArtist").innerHTML = artist;
    node.querySelector("#songCellDuration").innerHTML = readableSeconds(
      duration
    );
    return node;
  };

  var processResponse = function(obj) {
    const { type, data } = obj;
    log({ type, data });

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

        updateActiveSong(data.activeSong);
        triggerResize();
        break;
      case "searchResult":
        e("searchResult").innerHTML = ""; // delete old search result
        // debug({ gFiles: gPlaylistFiles });
        if (data.truncated) {
          showError("search result limited to " + data.maxResults);
        }
        if (data.searchResult.length == 0) {
          showError("nothing found");
        }
        data.searchResult.map(function(entry, i) {
          const { file } = entry;
          const node = newSongNode("searchEntry", entry, i);
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
        triggerResize();

        break;
      case "directoryList":
        e("directoryList").innerHTML = "";
        if (data.hasParent) {
          node = e("directoryListEntry").cloneNode(true);
          node.id = "dlRowParent";
          node.title = data.parent;
          node.classList.remove("hide");
          node.querySelector("#directoryName").innerHTML =
            data.parent != "" ? data.parent : "..";
          ["#directoryName", "#dlBrowse"].map(function(v) {
            node.querySelector(v).onclick = btnCommand("browse", data.parent);
          });
          e("directoryList").append(node);
        }
        data.directoryList.map(function(entry, i) {
          var node;
          if (entry.type == "directory") {
            node = e("directoryListEntry").cloneNode(true);
            node.id = "dlRow" + i;
            node.classList.remove("hide");
            node.querySelector("#directoryName").innerHTML = entry.directory;
            ["#directoryName", "#dlBrowse"].map(function(v) {
              node.querySelector(v).onclick = btnCommand(
                "browse",
                entry.directory
              );
            });
          } else {
            node = newSongNode("searchEntry", entry, i);
            const btn = node.querySelector("#srAdd");
            // disable button for files already in the playlist
            if (gPlaylistFiles.includes(entry.file)) {
              // FIXME: should we add a button to remove it from the playlist?
              btn.disabled = "disabled";
            }
            btn.onclick = function() {
              btn.disabled = "disabled";
              return command("add", entry.file);
            };
          }
          e("directoryList").append(node);
        });
        break;
    }
  };
  var openWebSocket = function() {
    ws = new WebSocket(ws_addr);
    ws.onopen = function(evt) {
      debug("OPEN");
      hideError();
    };
    ws.onclose = function(evt) {
      debug("CLOSE");
      ws = null;
      showError("no connection");
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
    debug("stop");
    ["play"].map(show);
    ["pause", "resume"].map(hide);
    ["stop", "next", "previous"].map(disable);
  };
  var play = function() {
    debug("play");
    ["pause"].map(show);
    ["resume", "play"].map(hide);
    ["stop", "next", "previous"].map(enable);
  };
  var pause = function() {
    debug("pause");
    ["resume"].map(show);
    ["pause", "play"].map(hide);
    ["stop", "next", "previous"].map(enable);
  };
  var togglePlayPause = function(state) {
    // debug(`togglePlayPause(${state})`);
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

  // add onclick function for all controls
  ["play", "resume", "pause", "stop", "next", "previous"].map(function(value) {
    e(value).onclick = function(evt) {
      debug(`Control: ${value}`);
      return command(value, "");
    };
  });

  // show the viewPlaylist
  show("viewPlaylist");
});

// eof