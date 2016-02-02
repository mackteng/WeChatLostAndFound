// Sends an XML encoded string to the specified URL
// and returns response
function sendXMLMessage(data, beforeSend, success){

	$.ajax({
  		type: "POST",
  		url:"http://ec2-52-0-73-218.compute-1.amazonaws.com",
  		data: data,
  		success: success,
		contentType: "text/xml",
		beforeSend: beforeSend,
		timeout:5000
	});
}




var BlockFinder = function(){





}

var DeleteFinder = function(){





}

var ChangeToActiveTag = function(){





}


var DeleteTag = function(){





};

var sendRegister = function(tagID, name, desc){

	$("#confirm_register").off('click');
	$("register_form").hide();

	var beforeSend = function(){
		wx.closeWindow();
	};

	var success = function(response){
	};

	var data = 
    "<xml><ToUserName><![CDATA[" + "gh_6df161a83822" + "]]></ToUserName>" +
    "<FromUserName><![CDATA[" + OpenID + "]]></FromUserName>" +
    "<CreateTime>" + new Date().getTime() + "</CreateTime>" +
    "<MsgType><![CDATA[event]]></MsgType>" + 
    "<Event><![CDATA[scancode_waitmsg]]></Event>" + 
    "<EventKey><![CDATA[RegisterTag]]></EventKey>" + 
    "<ScanCodeInfo><ScanType><![CDATA[qrcode]]></ScanType>"+
    "<ScanResult><![CDATA[" + tagID +"]]></ScanResult>"+
    "</ScanCodeInfo>"+
    "<ItemInfo><Name><![CDATA[" + name +  "]]></Name>"+
    "<Description><![CDATA[" + desc +"]]></Description>"+
    "</ItemInfo>"+
    "</xml>";

	sendXMLMessage(data, beforeSend, success);
}

// Opens the QR Code Scanner and shows the Item Description Form Upon Success
var openQR = function(){
	wx.scanQRCode({
    		needResult: 1, // 默认为0，扫描结果由微信处理，1则直接返回扫描结果，
    		scanType: ["qrCode"], // 可以指定扫二维码还是一维码，默认二者都有
    		success: function (res) {
			showMenu(res.resultStr);
		}
	});
}

// Shows the menu and then 
var showMenu = function(tagid) {
  alert(tagid);
  var $form = $("#register_form");
  $form.show();
  
  $form.find("#confirm_register").on('click', function() {
    
    var $name = $("#item_name").val();
    var $desc = $("#item_desc").val();
    sendRegister(tagid,$name,$desc);

  });

  $form.find("#cancel_register").on('click', function() {
    $form.hide();
  });
};

$(wx).ready(function(){
		$("#register_button").on('click', openQR);
 	}
);
