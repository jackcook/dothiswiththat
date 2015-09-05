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
  document.getElementsByTagName("iframe")[0].setAttribute("src", link);
});
