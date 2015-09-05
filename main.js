$(function() {
  var actionTags = [
    "read a book",
    "watch a movie"
  ];

  var languageTags = [
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

var go = document.getElementById("go");
go.addEventListener("click", function() {
  console.log("yey");
});
