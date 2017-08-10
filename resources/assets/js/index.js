
$(document).ready(function() { 
  todoList.forEach( function(element){addToList(element)})
  completedList.reverse().forEach( function(element){addToList(element)})
  $("#addTodo").submit(function(evt){addToDo(evt)})
})
      
var tasks = [];
var oldDate = Date.parse("March 21, 2012"); //just an old date

function prepareFormFromElement(formId) {

  form = document.getElementById(formId);
  formData = new FormData(form);

  var formBody = [];
  for(var pair of formData.entries()) {
    var encodedKey = encodeURIComponent(pair[0]);
    var encodedValue = encodeURIComponent(pair[1]);
    formBody.push(encodedKey + "=" + encodedValue);
  }

  return formBody.join("&");
}

function prepareFormFromObject(Object) {

  formData = new FormData(form);

  var formBody = [];

  for (var property in Object) {
    var encodedKey = encodeURIComponent(property);
    var encodedValue = encodeURIComponent(details[property]);
    formBody.push(encodedKey + "=" + encodedValue);
  }

  return formBody.join("&");
}

function addToDo(evt){
  evt.stopImmediatePropagation();
  var url = "/todo";

  formBody = prepareFormFromElement("addTodo");

  fetch( url, {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/x-www-form-urlencoded'
    },
    credentials: 'include', 
    body: formBody
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
  $("#myInput").val("");
}

function addToList(json) {
  var type = "append"
  var taskClass = "";
  var disabled = "";

  var completedon = Date.parse("March 21, 2010"); // to make sure we don't reuse old values
  var completedon = Date.parse(json.CompletedOn);
  
  if (completedon > oldDate) {
    type = "prepend"
    taskClass = "success"
    disabled = "disabled"
  }

  tasks[json.ID] = json;
  data = '\
          <tr id="task'+ json.ID +'" class="'+ taskClass +'">\
            <td>' + json.Title +'</td>\
            <td>\
              <div class="btn-group">\
                <button type="button" id="editTask'+ json.ID +'" class="btn btn-default" onclick="editTaskForm('+ json.ID +')">Edit</button>\
                <button type="button" id="completeTask'+ json.ID +'" class="btn btn-success" onclick="completeTask('+ json.ID +')" '+ disabled +'>Done</button>\
                <button type="button" id="deleteTask'+ json.ID +'" class="btn btn-danger" onclick="deleteTask('+ json.ID +')">Delete</button>\
              </div></td>\
          </tr>';
  
  if (type === 'prepend') {
    $("#completedList").prepend(data)
  } else {
    $("#todoList").append(data)
  }

}

function editTaskForm(jsonID){

  var json = tasks[jsonID]
  console.log(json)
  $('#editTaskModal').modal('show');
  $('#title').val(json.Title);
  $('#priority').val(json.Priority);
  if(json.Description.Valid === true) {
    $('#description').val(json.Description.String);
  }
}

function completeTask(taskId){

  var url = "/todo/"+ taskId;

  fetch( url, {
      method: 'PUT',
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded'
      },
      credentials: 'include', 
      body: "completed_on=now" 
  }).then(function(response) {
      return response.json()
  }).then(function(json){
      if (json.Status != 'success') {
        alert('An error occured, please try again');
      } else {
        // set a timeout to the alert goes away
        $("#task" + taskId).addClass("success");
        $("#completeTask" + taskId).prop("disabled", true);
        timeout = setTimeout(function() {
          $("#task" + taskId).remove();
          addToList(json.Data);
        }, 2000);
      }
  }).catch(function(err){
      alert('An error occured, please try again');
      console.log(err)
  });
}

function deleteTask(taskId){

  var url = "/todo/"+ taskId;

  fetch( url, {
      method: 'DELETE',
      credentials: 'include', 
  }).then(function(response) {
      return response.json()
  }).then(function(json){
      if (json.Status != 'success') {
        alert('An error occured, please try again');
      } else {
        $("#task" + taskId).addClass("danger");
        $("#deleteTask" + taskId).prop("disabled", true);
        $("#editTask" + taskId).prop("disabled", true);
        $("#completeTask" + taskId).prop("disabled", true);

        // set a timeout to the alert goes away
        timeout = setTimeout(function() {
          $("#task" + taskId).remove();
        }, 2000);
      }
  }).catch(function(err){
      alert('An error occured, please try again');
      console.log(err)
  });
}
