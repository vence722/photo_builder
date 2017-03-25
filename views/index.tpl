<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	
	<!-- jQuery -->
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
	
	<!-- bootstrap -->
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/css/bootstrap-theme.min.css" integrity="sha384-rHyoN1iRsVXV4nD0JutlnGaslCJuC7uwjduW9SVrLvRYooPp2bWYgmgJQIXwl/Sp" crossorigin="anonymous">
	<script src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
	
	<!-- waiting dialog -->
	<script src="/static/js/waiting_dialog.js"></script>

	<style type="text/css">
		.camera_name {
			font-size: 24px
		}
		
		.photo_list {
			min-height: 200px
		}
	</style>
	
	<script type="text/javascript">
		$(function () {
			var cameras = $(".camera_name");
			for (var i=0;i<cameras.size();i++) {
				var cid = cameras.eq(i).text();
				getPhotos(cid);
			}
		});
		
		function getPhotos(cid) {
			$.post("/", {cameraId: cid}, function(resp) {
				for (var j=0;j<resp.length;j++) {
					var photo = resp[j];
					var img_elem = $("<img src=\"data:image/jpeg;base64," + photo.DataBase64 + "\" style=\"max-height:200px\"/>");
					$("#photo_list_" + cid).append(img_elem);
				}
			});
		}
		
		function template1() {
			$("#loadingModal").modal("show");
			waitingDialog.show();
			$.post("/proess", {tmplID: "1"}, function (resp) {
				var photo = resp;
				var img_elem = $("<img src=\"data:image/jpeg;base64," + photo.DataBase64 + "\" style=\"max-height:500px\"/>");
				$("#result_div").html(img_elem);
				$("#loadingModal").modal("hide");
				waitingDialog.hide();
			});
		}
		
		function template2() {
			$("#loadingModal").modal("show");
			waitingDialog.show();
			$.post("/proess", {tmplID: "2"}, function (resp) {
				var photo = resp;
				var img_elem = $("<img src=\"data:image/jpeg;base64," + photo.DataBase64 + "\" style=\"max-height:500px\"/>");
				$("#result_div").html(img_elem);
				$("#loadingModal").modal("hide");
				waitingDialog.hide();
			});
		}
	</script>
</head>

<body>
	<div class="modal hide" id="loadingModal">
	</div>
	{{with .Cameras}}
		{{range .}}
			<div style="width:360px;float:left;border:1px solid black;text-align:center">
				<div class="camera_name">{{.}}</div>
				<div id="photo_list_{{.}}" class="photo_list"></div>
			</div>
		{{end}}
	{{end}}
	<br/>
	<input type="button" class="btn btn-success" style="width:48%" value="Template1" onclick="template1()"/>
	<input type="button" class="btn btn-danger" style="width:48%" value="Template2" onclick="template2()"/>
	<div id="result_div" style="margin-top:50px; text-align: center;"></div>
</body>