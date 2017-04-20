<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="utf-8">
	
	<!-- jQuery -->
	<script src="https://ajax.googleapis.com/ajax/libs/jquery/1.11.2/jquery.min.js"></script>
	
	<!-- Lodash -->
	<script src="https://cdn.jsdelivr.net/lodash/4.17.4/lodash.min.js"></script>

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
		var selected = [];
		var initCount = 0;

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
					var img_elem = $("<img cid=\"" + cid + "\" filename=\"" + photo.FileName + "\" src=\"data:image/jpeg;base64," 
						+ photo.DataBase64 + "\" style=\"width:100%;max-height:100px;\" onclick=\"choosePhoto(" + j + ")\"/>");
					$("#photo_list_" + cid).append(img_elem);
				}
				if (++initCount === $(".camera_name").size()) {
					choosePhoto(0);
				}
			});
		}

		function choosePhoto(index) {
			for (var i in selected) {
				var sel = selected[i];
				$("img[cid='" + sel.cid +  "'][filename='" + sel.filename + "']").css("border", "none");
			}
			selected = [];
			var cameras = $(".camera_name");
			for (var i=0;i<cameras.size();i++) {
				var cameraId = cameras.eq(i).text();
				var imgs = $("#photo_list_" + cameraId + " > img");
				var cid = imgs.eq(index).attr("cid");
				var filename = imgs.eq(index).attr("filename");
				if (cid && filename) {
					selected.push({cid, filename});
					imgs.eq(index).css("border", "5px solid red");
				}
			}
		}
		
		function template1() {
			if (!validateChoose()) {
				return;
			}
			$("#result_div").html("");
			$("#loadingModal").modal("show");
			waitingDialog.show();
			$.post("/proess", {tmplID: "1", "selected": JSON.stringify(_.compact(selected))}, function (resp) {
				var photos = resp;
				var photo = resp[1];
				var img_elem = $("<img src=\"data:image/jpeg;base64," + photo.DataBase64 + "\" style=\"max-height:500px\"/>");
				$("#result_div").append(img_elem);
				$("#result_div").append("<br/><br/>");
				photo = resp[0];
				img_elem = $("<img src=\"data:image/jpeg;base64," + photo.DataBase64 + "\" style=\"max-height:500px\"/>");
				$("#result_div").append(img_elem);
				$("#loadingModal").modal("hide");
				waitingDialog.hide();
			});
		}
		
		function template2() {
			if (!validateChoose()) {
				return;
			}
			$("#result_div").html("");
			$("#loadingModal").modal("show");
			waitingDialog.show();
			$.post("/proess", {tmplID: "2", "selected": JSON.stringify(_.compact(selected))}, function (resp) {
				var photo = resp;
				var photo = resp[1];
				var img_elem = $("<img src=\"data:image/jpeg;base64," + photo.DataBase64 + "\" style=\"max-height:500px\"/>");
				$("#result_div").append(img_elem);
				$("#result_div").append("<br/><br/>");
				photo = resp[0];
				img_elem = $("<img src=\"data:image/jpeg;base64," + photo.DataBase64 + "\" style=\"max-height:500px\"/>");
				$("#result_div").append(img_elem);
				$("#loadingModal").modal("hide");
				waitingDialog.hide();
			});
		}

		function validateChoose() {
			var cameras = $(".camera_name");
			if (_.compact(selected).length < cameras.size()) {
				$("#dialog").modal();
				return false;
			}
			return true;
		}
	</script>
</head>

<body>
	<div class="modal hide" id="loadingModal">
	</div>
	<div class="modal fade" id="dialog" role="dialog">
		<div class="modal-dialog modal-sm">
		<div class="modal-content">
			<div class="modal-header">
			<button type="button" class="close" data-dismiss="modal">&times;</button>
			<h4 class="modal-title">Tips</h4>
			</div>
			<div class="modal-body">
			<p>Please choose at least 1 photo for each camera!</p>
			</div>
			<div class="modal-footer">
			<button type="button" class="btn btn-default" data-dismiss="modal">Close</button>
			</div>
		</div>
		</div>
	</div>
	{{with .Cameras}}
		{{range .}}
			<div style="width:16%;float:left;border:1px solid black;text-align:center">
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