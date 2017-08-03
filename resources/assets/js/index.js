function addToDo(evt){
  evt.stopImmediatePropagation();
  var url = "/add-todo?title="+ $("#myInput").val();
  $("#myInput").val("");

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
  var type = "append"
  var taskClass = "";
  var disabled = "";

  var oldDate = Date.parse("March 21, 2012"); //just an old date
  var completedon = Date.parse("March 21, 2010"); // to make sure we don't reuse old values
  completedon = Date.parse(json.CompletedOn);
  
  if (completedon > oldDate) {
    type = "prepend"
    taskClass = "success"
    disabled = "disabled"
  }

                //<button type="button" id="editTask'+ json.ID +'" class="btn btn-default" onclick="editTask('+ json.ID +')">Edit</button>\
  data = '\
          <tr id="task'+ json.ID +'" class="'+ taskClass +'">\
            <td>' + json.Title +'</td>\
            <td>\
              <div class="btn-group">\
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

function completeTask(taskId){

  var url = "/complete-todo/"+ taskId;

  fetch( url, {
      method: 'get',
      credentials: 'include', 
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

  var url = "/delete-todo/"+ taskId;

  fetch( url, {
      method: 'get',
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
