var textarea = $('textarea').hide();
var editor = ace.edit("editor");

editor.setTheme("ace/theme/github");
editor.getSession().setMode("ace/mode/markdown");
editor.getSession().setOption("wrap", true);

editor.getSession().on('change', function(){
  textarea.val(editor.getSession().getValue());
});
