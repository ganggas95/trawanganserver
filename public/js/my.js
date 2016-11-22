function readURL(input) {
	if (input.files && input.files[0]) {
		var reader = new FileReader();

		reader.onload = function (e) {
			$('#blah').attr('src', e.target.result);
		}
		reader.readAsDataURL(input.files[0]);
	}
}
$("#uploader").change(function(){
	readURL(this);
});

$(document).on("click", ".delete-dialog", function () {
	var serviceId = $(this).data('id');
	$(".modal-body #idService").val( serviceId );
});

var counter = 1;
var limit = 3;

function addInput(divName) {
    if (counter == limit) {
        alert("You have reached the limit of adding " + counter + " inputs");
    } else {
        var newInput = document.createElement('input');
        var newLabel = document.createElement('label');
        newInput.type = "file"
        newInput.className = "form-control"
        newInput.tagName = "foto2";
        newInput.id = "foto2";
        newLabel.htmlFor = "foto2";
        newLabel.innerHTML = "Foto layanan2";
        document.getElementById(divName).appendChild(newInput);
        document.getElementById(divName).appendChild(newLabel);
        counter++;
    }
}