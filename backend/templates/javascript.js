var app = new Vue({
  el: "#web-mpc",
  delimiters: ["${", "}"], // do not clash with gos templates
  data: {
    activeSong: {},
    elapsed: 0,
    duration: 0,
    progress: 0,
    state: {},
    view: "queue",
    views: [
      { name: "queue", class: "ctrlViewList", icon: "#icon-queue" },
      { name: "search", class: "ctrlViewSearch", icon: "#icon-search" },
      { name: "directory", class: "ctrlViewBrowse", icon: "#icon-folder" },
      { name: "playlists", class: "ctrlViewPlaylists", icon: "#icon-list" }
    ],
    modes: [
      { name: "random", class: "ctrlModeRandom" },
      { name: "repeat", class: "ctrlModeRepeat" },
      { name: "single", class: "ctrlModeSingle" },
      { name: "consume", class: "ctrlModeConsume" }
    ],
    queue: {},
    queuePage: 1,
    queueLastPage: 1,
    queueAction: [
      { cmd: "play", class: "songPlay", icon: "#icon-play" },
      { cmd: "remove", class: "songRemove", icon: "#icon-ban" }
    ],
    searchResult: {},
    searchText: "",
    searchPage: 1,
    searchLastPage: 1,
    priorities: [
      { value: 255, class: "songPrioUp", icon: "#icon-top" },
      { value: 127, class: "songPrioMiddle", icon: "#icon-middle" },
      { value: 0, class: "songPrioDown", icon: "#icon-bottom" }
    ]
  },
  methods: {
    toggleMode: function(mode) {
      if (this.$data.state[mode]) {
        command(mode, "disable");
      } else {
        command(mode, "enable");
      }
    },
    modeIcon: function(mode) {
      if (this.$data.state[mode]) {
        return "#icon-" + mode; // icon-random
      } else {
        return "#icon-" + mode + "-disabled"; // icon-random-disabled
      }
    },
    command: function(cmd, data) {
      return command(cmd, data.toString());
    },
    search: function(cmd, value, entry) {
      switch (cmd) {
        case "add":
          entry.queued = true;
          return command("add".entry.file);
        case "addPrio":
          entry.queued = true;
          return command(cmd, value + ":" + entry.file);
      }
    },
    readableSeconds: function(value) {
      var min = parseInt(value / 60);
      var sec = parseInt(value % 60);
      if (sec < 10) {
        sec = "0" + sec;
      }
      return min + ":" + sec;
    }
  }
});

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
var addcls = function(el, cls) {
  el.classList.add(cls);
};
// short for document.getElementById(id).classList.rm(class)
var rmcls = function(el, cls) {
  el.classList.remove(cls);
};
var hide = function(el) {
  addcls(el, "hide");
};
var show = function(el) {
  rmcls(el, "hide");
};
var disable = function(el) {
  el.disabled = "disabled";
};
var enable = function(el) {
  el.disabled = "";
};
var hideId = function(id) {
  hide(e(id));
};
var showId = function(id) {
  show(e(id));
};
var disableId = function(id) {
  disable(e(id));
};
var enableId = function(id) {
  enable(e(id));
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
  while (fs >= minFS && getTextWidth(txt, fs) + fs > maxWidth) {
    fs -= 2;
  }
  if (fs < minFS) {
    fs = minFS;
    maxWidth *= 2; // we allow one page break
    var truncated = false;
    while (txt.length > 8 && getTextWidth(txt, fs) + fs > maxWidth) {
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
    resize(el[i], 12, 24); // FIXME: take body font-size as max
  }
};

var triggerResize = function() {
  window.dispatchEvent(new Event("resize")); // trigger resize events on the song stuff
};

var ws_addr, ws;

// send a command on the websocket
var command = function(cmd, data) {
  if (!ws) return false;
  log({ cmd, data });
  if (cmd == "search" && data.length < 3) {
    return false;
  }
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
var gErrors = [];
var gPlaylistFiles = [];
var readableSeconds = function(value) {
  var min = parseInt(value / 60);
  var sec = parseInt(value % 60);
  if (sec < 10) {
    sec = "0" + sec;
  }
  return min + ":" + sec;
};
var gPage = 1;

// this functions runs forever and gets called every second to update the
// elapsed and duration information of the active song
var updateProgress = function() {
  if (app.state["play"] == "play" && app.elapsed < app.duration) {
    // increment the seconds if playing and not finished
    app.elapsed += 1;
  }
  if (app.state["play"] == "stop") {
    app.progress = 0;
  } else {
    app.progress = (app.elapsed / app.duration) * 100;
  }
  setTimeout(updateProgress, 1000);
};

var showError = function(msg) {
  const el = e("mainError");
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
  el.innerHTML = str;
  show(el);

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
  const el = e("mainError");
  el.innerHTML = "";
  hide(el);
};

var updateStatus = function(data) {
  debug("updateStatus");
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
  app.state = {
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
      app.duration = duration;
      app.elapsed = elapsed;
      break;
    case "stop":
      stop();
      app.duration = 0;
      app.elapsed = 0;
      break;
  }
};

var updateActiveSong = function(data) {
  const { file, artist, title, album_artist, album, position } = data;
  e("ctrlSong").title = file;
  app.activeSong = data;
  const el = e("remove");
  if (position == -1) {
    disable(el);
  } else {
    enable(el);
    el.onclick = btnCommand("remove", position.toString());
  }
};

var newSongNode = function(id, entry, nr) {
  const {
    file,
    disc,
    track,
    artist,
    title,
    album,
    album_artist,
    duration
  } = entry;
  var node = e(id).cloneNode(true);
  node.classList.remove("hide");
  node.title = file;
  node.id = id + nr;
  if (disc) {
    node.querySelector("#songRowDisc").innerHTML = disc;
  }
  if (track) {
    node.querySelector("#songRowTrack").innerHTML = track;
  }
  node.querySelector("#songRowTitle").innerHTML = title;
  node.querySelector("#songRowArtist").innerHTML = artist;
  // only show album artist if it's different
  node.querySelector("#songRowAlbum").innerHTML =
    album + (artist != album_artist ? "<wbr>[" + album_artist + "]" : "");
  node.querySelector("#songRowDuration").innerHTML = readableSeconds(duration);
  return node;
};

var pagination = function(page, lastPage, cmdStr) {
  if (lastPage == 1) {
    return null;
  }
  var node = e("pagination").cloneNode(true);
  node.classList.remove("hide");
  previousPage = page - 1;
  nextPage = page + 1;
  node.querySelector("#currentPage").innerHTML = page + "/" + lastPage;
  if (page == 1) {
    disable(node.querySelector("#firstPage"));
  } else {
    enable(node.querySelector("#firstPage"));
    node.querySelector("#firstPage").onclick = btnCommand(cmdStr, "1");
  }
  if (previousPage < 1) {
    disable(node.querySelector("#previousPage"));
  } else {
    enable(node.querySelector("#previousPage"));
    node.querySelector("#previousPage").onclick = btnCommand(
      cmdStr,
      previousPage.toString()
    );
  }
  if (nextPage > lastPage) {
    disable(node.querySelector("#nextPage"));
  } else {
    enable(node.querySelector("#nextPage"));
    node.querySelector("#nextPage").onclick = btnCommand(
      cmdStr,
      nextPage.toString()
    );
  }
  if (page == lastPage) {
    disable(node.querySelector("#lastPage"));
  } else {
    enable(node.querySelector("#lastPage"));
    node.querySelector("#lastPage").onclick = btnCommand(
      cmdStr,
      lastPage.toString()
    );
  }
  return node;
};

var processResponse = function(obj) {
  const { type, data } = obj;
  log({ type, data });

  switch (type) {
    case "error":
    case "info":
      showError(data);
      break;
    case "version":
      log("MPD protocol version: " + data);
      break;
    case "update":
      updateStatus(data.status);
      app.queue = data.queue;
      app.queuePage = data.page;
      app.queueLastPage = data.lastPage;
      if (false) {
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
          if (isNext) {
            addcls(node, "nextSong");
          }
          // FIXME: this only makes sense in certain modes (random + ?)
          const btnPrio1 = node.querySelector("#plPrio1");
          const btnPrio2 = node.querySelector("#plPrio2");
          const btnPrio3 = node.querySelector("#plPrio3");
          if (app.state["random"]) {
            // see explanation above
            if (
              (state != "stop" || activeFile != "") &&
              (isActive || queueSize <= 2)
            ) {
              disable(btnPrio1);
              disable(btnPrio2);
              disable(btnPrio3);
            } else {
              switch (prio) {
                case 0:
                  disable(btnPrio3);
                  break;
                case 127:
                  disable(btnPrio2);
                  break;
                case 255:
                  disable(btnPrio1);
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
      }
      updateActiveSong(data.activeSong);
      triggerResize();
      break;

    case "searchResult":
      const { searchResult, page, lastPage } = data;
      app.searchResult = searchResult;
      app.searchPage = page;
      app.searchLastPage = lastPage;
      triggerResize();

      break;
    case "directoryList":
      e("directoryList").innerHTML = "";
      if (data.hasParent) {
        node = e("directoryListEntry").cloneNode(true);
        node.id = "dlRowParent";
        node.title = data.parent;
        show(node);
        node.querySelector("#directoryName").innerHTML =
          data.parent != "" ? data.parent : "..";

        ["#directoryName", "#browse"].map(function(v) {
          node.querySelector(v).onclick = btnCommand("browse", data.parent);
        });
        disable(node.querySelector("#add"));
        e("directoryList").append(node);
      }
      data.directoryList.map(function(entry, i) {
        const { type, directory, file } = entry;
        var node;
        if (type == "directory") {
          node = e("directoryListEntry").cloneNode(true);
          node.id = "dlRow" + i;
          rmcls(node, "hide");
          node.querySelector("#directoryName").innerHTML = directory;
          ["#directoryName", "#browse"].map(function(v) {
            node.querySelector(v).onclick = btnCommand("browse", directory);
          });
          const btn = node.querySelector("#add");
          btn.onclick = function() {
            disable(btn);
            return command("add", directory);
          };
        } else {
          node = newSongNode("searchEntry", entry, i);
          const btn = node.querySelector("#srAdd");
          // disable button for files already in the playlist
          if (gPlaylistFiles.includes(file)) {
            // FIXME: should we add a button to remove it from the playlist?
            disable(btn);
          }
          btn.onclick = function() {
            disable(btn);
            return command("add", file);
          };
        }
        e("directoryList").append(node);
        triggerResize();
      });
      break;
    case "playlistList":
      e("playlistList").innerHTML = "";
      if (data.hasParent) {
        node = e("playlistListEntry").cloneNode(true);
        node.id = "dlRowParent";
        node.title = data.parent;
        show(node);
        node.querySelector("#playlistName").innerHTML =
          data.parent != "" ? data.parent : "..";

        ["#playlistName", "#plBrowse"].map(function(v) {
          node.querySelector(v).onclick = btnCommand("browse", data.parent);
        });
        disable(node.querySelector("#plAdd"));
        e("playlistList").append(node);
      }
      data.playlistList.map(function(entry, i) {
        const { playlist } = entry;
        var node = e("playlistListEntry").cloneNode(true);
        node.id = "dlRow" + i;
        rmcls(node, "hide");
        node.querySelector("#playlistName").innerHTML = playlist;
        ["#playlistName", "#browse"].map(function(v) {
          node.querySelector(v).onclick = btnCommand("browse", playlist);
        });
        const btn = node.querySelector("#add");
        btn.onclick = function() {
          disable(btn);
          return command("add", playlist);
        };
        e("playlistList").append(node);
        triggerResize();
      });
      break;
  }
};

var shouldRefresh = false;
var openWebSocket = function() {
  ws = new WebSocket(ws_addr);
  ws.onopen = function(evt) {
    debug("OPEN");
    hideError();
    if (shouldRefresh) {
      location.reload(); // force refresh to get latest html etc
    }
  };
  ws.onclose = function(evt) {
    debug("CLOSE");
    ws = null;
    showError("no connection&hellip; reconnecting in 4 seconds&hellip;");
    shouldRefresh = true;
    setTimeout(openWebSocket, 4000);
  };
  ws.onmessage = function(evt) {
    processResponse(JSON.parse(evt.data));
  };
  ws.onerror = function(evt) {
    log({ evt });
  };
};

// FIXME: play / pause / stop doesn't work as expected (does not enable buttons all the time)
var stop = function() {
  debug("stop");
  ["pause", "resume", "stop", "next", "previous"].map(hideId);
  ["play", "remove"].map(showId);
  ["play", "remove"].map(enableId);
};
var play = function() {
  debug("play");
  ["resume", "play", "remove"].map(hideId);
  ["pause", "stop", "next", "previous"].map(showId);
  ["pause", "stop", "next", "previous"].map(enableId);
};
var pause = function() {
  debug("pause");
  ["pause", "play", "remove"].map(hideId);
  ["resume", "stop", "next", "previous"].map(showId);
  ["resume", "stop", "next", "previous"].map(enableId);
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

window.addEventListener("load", function(evt) {
  ws_addr = e("ws").value; // read from hidden input field
  addEvent(window, "resize", resizer);
  addEvent(window, "orientationchange", resizer);

  openWebSocket();

  updateProgress();
  // this often leads to sending the search twice, disabled until I find time
  // to fix it
  /*
  e("searchText").onchange = function(evt) {
    return command("search", e("searchText").value);
  };
  */

  ["play", "resume", "pause", "stop", "next", "previous", "clean"].map(function(
    value
  ) {
    e(value).onclick = btnCommand(value, "");
  });

  window.onfocus = function(event) {
    // request a fresh status as some browsers (e. g. Chrome) suspend our
    // progress setTimeout functions
    command("updateRequest", gPage.toString());
  };
});

// eof
