function get_link(action, language, callback) {
  file = "languages/" + language + ".json";

  request = new XMLHttpRequest();
  request.onreadystatechange = function() {
    if (request.readyState == 3) {
      var response = JSON.parse(request.responseText)[action];
      var link = response[Math.floor(Math.random() * response.length)];
      callback(link);
    }
  }

  request.open("GET", file, true);
  request.send();
}

var action = "";
var lang = "";

var paramslist = window.location.href.split("?")[1].split("&");

for (var i = 0; i < paramslist.length; i++) {
  var param = paramslist[i];
  var key = param.split("=")[0];
  var value = param.split("=")[1];

  if (key == "lang") {
    lang = value;
  } else if (key == "action") {
    action = value;
  }
}

get_link(action, lang, function(link) {
  var type = link.split(":")[0];
  var value = link.split(":")[1];
  var url = "nothing.html";

  if (type == "gplaytv") {
    url = "https://play.google.com/store/tv/show?id=" + value;
  } else if (type == "http") {
    url = link;
  } else if (type == "soundcloud") {
    url = "https://w.soundcloud.com/player/?url=https%3A//api.soundcloud.com/tracks/" + value + "&amp;auto_play=false&amp;hide_related=false&amp;show_comments=true&amp;show_user=true&amp;show_reposts=false&amp;visual=true";
  } else if (type == "youtube") {
    url = "https://youtube.com/embed/" + value;
  }

  document.getElementsByTagName("iframe")[0].setAttribute("src", url);
});
