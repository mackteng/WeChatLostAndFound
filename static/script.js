// encodes Object to an XML string
// and returns
function encodeObject(obj){






}

// Sends an XML encoded string to the specified URL
// and returns response
function sendXMLMessage(string, url){






}
var BlockFinder = function(){





}

var DeleteFinder = function(){





}

var ActivateTag = function(){





}


var DeleteTag = function(){





}



var sendRegister = function(tagid){
	alert(OpenID);
	alert(tagid);
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
  var $form = $("#register_form");
  $form.show();
  
  $form.find("#confirm_register").on('click', function() {
    sendRegister(tagid);
    $form.hide();
  });

  $form.find("#cancel_register").on('click', function() {
    $form.hide();
  });
};

$(wx).ready(function(){
		$("#register_button").on('click', openQR);
 	}
);
