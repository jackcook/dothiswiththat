$(function() {
  var actionTags = [
    "read a book",
    "watch a movie"
  ];

  var languageTags = [
    "English",
    "French",
    "Spanish",
    "Swedish"
  ];

  $("#actions").autocomplete({
    source: actionTags
  });

  $("#languages").autocomplete({
    source: languageTags
  });
});

var actions = {
  "read a book": "book",
  "watch a movie": "movie",
  "read the news": "news"
}

var languages = {
  "english": "en",
  "spanish": "es",
  "french": "fr",
  "swedish": "sv"
}

function get_links(action, language) {
  file = "languages/" + languages[language.toLowerCase()] + ".json";

  request = new XMLHttpRequest();
  request.onreadystatechange = function() {
    if (request.readyState == 3) {
      var response = JSON.parse(request.responseText)[actions[action]];
      console.log(response);
    }
  }

  request.open("GET", file, true);
  request.send();
}

var go = document.getElementById("go");
var actionsElement = document.getElementById("actions");
var languagesElement = document.getElementById("languages");

go.addEventListener("click", function() {
  get_links(actionsElement.value, languagesElement.value);
});
