// Create a "close" button and append it to each list item
var myNodelist = document.getElementsByTagName("LI");
var i;
for (i = 0; i < myNodelist.length; i++) {
  var span = document.createElement("SPAN");
  var txt = document.createTextNode("\u00D7");
  span.className = "close";
  span.appendChild(txt);
  myNodelist[i].appendChild(span);
}

// Click on a close button to hide the current list item
var close = document.getElementsByClassName("close");
var i;
for (i = 0; i < close.length; i++) {
  close[i].onclick = function() {
    var div = this.parentElement;
    div.style.display = "none";
  }
}

// Add a "checked" symbol when clicking on a list item
// var list = document.querySelector('ul');
// list.addEventListener('click', function(ev) {
//   if (ev.target.tagName === 'LI') {
//     ev.target.classList.toggle('checked');
//   }
// }, false);

// Create a new list item when clicking on the "Add" button
function addToDo(evt){
  evt.stopImmediatePropagation();
  var url = "/add-todo?title="+ document.getElementById('myInput').value;
  var todoAddForm = new FormData(document.getElementById('addTodo'));

  fetch( url, {
      method: 'get',
      credentials: 'include', 
  }).then(function(response) {
      return response.json()
  }).then(function(json){
      if (json.Status != 'success') {
          alert('An error occured, please try again');
      } else addToList(json.Data)
  }).catch(function(err){
      alert('An error occured, please try again');
      console.log(err)
  });
}

function addToList(json) {
  $("#myInput").val("");
  $("#listBody").append('\
                        <tr id="task'+ json.ID +'">\
                          <td>' + json.Title +'</td>\
                          <td>\
                            <div class="btn-group">\
                              <button type="button" id="editTask'+ json.ID +'" class="btn btn-default" onclick="editTask('+ json.ID +')">Edit</button>\
                              <button type="button" id="completeTask'+ json.ID +'" class="btn btn-success" onclick="completeTask('+ json.ID +')">Done</button>\
                              <button type="button" id="deleteTask'+ json.ID +'" class="btn btn-danger" onclick="deleteTask('+ json.ID +')">Delete</button>\
                            </div></td>\
                        </tr>'
                    );

}

function completeTask(taskId){
  $("#task" + taskId).addClass("success");
  $("#completeTask" + taskId).prop("disabled", true);
}
