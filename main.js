String.prototype.capitalizeFirstLetter = function() {
    return this.charAt(0).toUpperCase() + this.slice(1);
}

var actions = {
  "read a book": "books",
  "watch a movie": "movies",
  "listen to music": "music",
  "read the news": "news",
  "watch a tv show": "tv"
}

var languages = {
  "Chinese": "zh",
  "English": "en",
  "French": "fr",
  "Italian": "it",
  "Japanese": "ja",
  "Korean": "ko",
  "Spanish": "es",
  "Swedish": "sv"
}

$("#actions").autocomplete({
  source: Object.keys(actions)
});

$("#languages").autocomplete({
  source: Object.keys(languages)
});

var go = document.getElementById("go");
var actionsElement = document.getElementById("actions");
var languagesElement = document.getElementById("languages");

go.addEventListener("click", function() {
  var action = actions[actionsElement.value.toLowerCase()];
  var lang = languages[languagesElement.value.toLowerCase().capitalizeFirstLetter()];

  file = "languages/" + lang + ".json";

  request = new XMLHttpRequest();
  request.onreadystatechange = function() {
    if (request.readyState == 3) {
      var response = JSON.parse(request.responseText)[action];

      if (action == "movies") {
        window.location = "movies?lang=" + lang;
      } else if (action == "music") {
        window.location = "music?lang=" + lang;
      } else if (action == "news") {
        window.location = "news?lang=" + lang;
      } else {
        var link = response[Math.floor(Math.random() * response.length)];

        var type = link.split(":")[0];
        var value = link.split(":")[1];
        var url = "nothing.html";

        if (type == "gplaytv") {
          url = "https://play.google.com/store/tv/show?id=" + value;
        } else if (type == "http" || type == "https") {
          url = link;
        } else if (type == "youtube") {
          url = "https://youtube.com/watch?v=" + value;
        }

        window.location = url;
      }
    }
  }

  request.open("GET", file, true);
  request.send();
});

var flag = 1;

window.setInterval(function() {
  var action = actionsElement.getAttribute("placeholder");
  var newAction = action;

  var language = languagesElement.getAttribute("placeholder");
  var newLanguage = language;

  while (action === newAction) {
    newAction = Object.keys(actions)[Math.floor(Math.random() * Object.keys(actions).length)];
  }

  while (language === newLanguage) {
    newLanguage = Object.keys(languages)[Math.floor(Math.random() * Object.keys(languages).length)];
  }

  if (flag == 1) {
    actionsElement.setAttribute("placeholder", newAction);
  } else {
    languagesElement.setAttribute("placeholder", newLanguage);
  }

  flag += flag == 1 ? -1 : 1;
}, 700);
