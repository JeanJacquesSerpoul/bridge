function fr(v) {
	if (v == "T") {
		return "10"
	}
	if (v == "J") {
		return "V"
	}
	if (v == "Q") {
		return "D"
	}
	if (v == "K") {
		return "R"
	}
	return v
}

function translateScore(s) {
	s = s.replace(/NS/gi, "xxx").trim();
	s = s.replace(/:S/gi, "yyy").trim();
	s = s.replace(/S/gi, "P").trim();
	s = s.replace(/D/gi, "K").trim();
	s = s.replace(/C/gi, "T").trim();
	s = s.replace(/H/gi, "C").trim();
	s = s.replace(/N/gi, "SA").trim();
	s = s.replace(/EW/gi, "EO").trim();
	s = s.replace(/xxx/gi, "NS").trim();
	s = s.replace(/yyy/gi, ":S").trim();
	return s
}

function trans(suit) {
	var c = "";
	var t = "";
	for (var i = 0; i < suit.length; i++) {
		t = fr(suit[i]);
		c = c + t + " ";
	}
	return c
}

function download(filename, text) {
	var pom = document.createElement('a');
	pom.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
	pom.setAttribute('download', filename);

	if (document.createEvent) {
		var event = document.createEvent('MouseEvents');
		event.initEvent('click', true, true);
		pom.dispatchEvent(event);
	} else {
		pom.click();
	}
}

function convertCards(card) {
	card = card.replace(/10/gi, "T").trim();
	card = card.replace(/R/gi, "K").trim();
	card = card.replace(/D/gi, "Q").trim();
	card = card.replace(/V/gi, "J").trim();
	card = card.replace(/A/gi, "A").trim();
	return card
}

function convertSuit(suit) {
	if (suit.trim() == "") {
		return "-1"
	}
	return suit
}

function convertPoints(points) {
	if (points.trim() == "") {
		return "-1"
	}
	return points
}


function getSuit() {
	var mask = ""
	mask += convertSuit($('#c-n-s').val()) + "," + convertSuit($('#c-n-h').val()) + "," + convertSuit($('#c-n-d').val()) + "," + convertSuit($('#c-n-c').val()) + ","
	mask += convertSuit($('#c-e-s').val()) + "," + convertSuit($('#c-e-h').val()) + "," + convertSuit($('#c-e-d').val()) + "," + convertSuit($('#c-e-c').val()) + ","
	mask += convertSuit($('#c-s-s').val()) + "," + convertSuit($('#c-s-h').val()) + "," + convertSuit($('#c-s-d').val()) + "," + convertSuit($('#c-s-c').val()) + ","
	mask += convertSuit($('#c-w-s').val()) + "," + convertSuit($('#c-w-h').val()) + "," + convertSuit($('#c-w-d').val()) + "," + convertSuit($('#c-w-c').val())
	return mask
}

function getPoints() {
	var mask = ""
	mask += convertPoints($('#nmin').val()) + "," + convertPoints($('#nmax').val()) + "," + convertPoints($('#emin').val()) + "," + convertPoints($('#emax').val()) + ","
	mask += convertPoints($('#smin').val()) + "," + convertPoints($('#smax').val()) + "," + convertPoints($('#wmin').val()) + "," + convertPoints($('#wmax').val())
	return mask
}


function getPbn() {
	var mask = ""
	mask = convertCards($('#n-s').val()) + "." + convertCards($('#n-h').val()) + "." + convertCards($('#n-d').val()) + "." + convertCards($('#n-c').val());
	mask += " ";
	mask += convertCards($('#e-s').val()) + "." + convertCards($('#e-h').val()) + "." + convertCards($('#e-d').val()) + "." + convertCards($('#e-c').val());
	mask += " ";
	mask += convertCards($('#s-s').val()) + "." + convertCards($('#s-h').val()) + "." + convertCards($('#s-d').val()) + "." + convertCards($('#s-c').val());
	mask += " ";
	mask += convertCards($('#w-s').val()) + "." + convertCards($('#w-h').val()) + "." + convertCards($('#w-d').val()) + "." + convertCards($('#w-c').val());
	return mask
}

function showDesk(r) {
	var hand = r.split(" ");
	var suit = hand[0].split(".");
	$('#north-spade').text(trans(suit[0]));
	$('#north-heart').text(trans(suit[1]));
	$('#north-diamond').text(trans(suit[2]));
	$('#north-club').text(trans(suit[3]));

	var suit = hand[1].split(".");
	$('#east-spade').text(trans(suit[0]));
	$('#east-heart').text(trans(suit[1]));
	$('#east-diamond').text(trans(suit[2]));
	$('#east-club').text(trans(suit[3]));

	var suit = hand[2].split(".");
	$('#south-spade').text(trans(suit[0]));
	$('#south-heart').text(trans(suit[1]));
	$('#south-diamond').text(trans(suit[2]));
	$('#south-club').text(trans(suit[3]));

	var suit = hand[3].split(".");
	$('#west-spade').text(trans(suit[0]));
	$('#west-heart').text(trans(suit[1]));
	$('#west-diamond').text(trans(suit[2]));
	$('#west-club').text(trans(suit[3]));
}

function getVersion() {
	$.ajax({
		url: 'api/version',
		type: 'GET',
		success: function (result) {
			$("#version").text("API version : " + result)
		}
	});
}

function waitHide(){
	$('#wait').hide();
	//$('#vuldeal').show();
	
}

function waitShow(){
	$('#wait').show();
	//$('#vuldeal').hide();
	
}

function getFromIndex() {
	var indexValue;
	indexValue = $('#index-donne').val();
	$('#getfromindexbtn').prop("disabled", true);
	initCalc();
	$.ajax({
		url: 'api/index',
		type: 'GET',
		data: {
			value: indexValue
		},
		success: function (result) {
			var result = jQuery.parseJSON(result);
			var r = result.pbn;
			$('#np').text("N:Points:" + result.np);
			$('#ep').text("E:Points:" + result.ep);
			$('#sp').text("S:Points:" + result.sp);
			$('#wp').text("O:Points:" + result.wp);
			$("#currentpbn").val(r);
			showDesk(r);
			setSuccess();
		},
		error: function (result, statut, error) {
			setError();
		},
		complete: function (resultat, statut) {
			waitHide();
		}
	});

}

function pbnOne() {
	if ($("#okpbn").text() == "Y") {
		var currentPbn = $("#currentpbn").val();
		var pbn = "";
		var dealer = $("#dealer").find('option:selected').val()
		var comment = $("#comment").val();
		var vulnerable = $("#vulnerable").find('option:selected').val()
		if (comment != "") {
			pbn = "% " + comment + "\n";
		}
		pbn = pbn + "% -Index: " + $("#index-donne").val() + "\n";
		pbn += "[Dealer \"" + dealer + "\"]\n[Vulnerable \"" + vulnerable + "\"]\n";
		result = pbn + "[Deal \"N:" + currentPbn + "\"]";
		download("donne.pbn", result)
	}
}

function pbn() {
	var count = $("#count").val();
	var dealer = $("#dealer").find('option:selected').val()
	var comment = $("#comment").val();
	var vulnerable = $("#vulnerable").find('option:selected').val()
	$('#pbnbtn').prop("disabled", true);
	initCalc();
	if (activeTabs() == 0) {
		var mask = getPbn();
		$.ajax({
			url: 'api/maskmultipbn',
			type: 'POST',
			data: {
				mask: mask,
				count: count,
				dealer: dealer,
				comment: comment,
				vulnerable: vulnerable
			},
			success: function (result) {
				setSuccess();
				download("donnes.pbn", result)
			},
			error: function (result, statut, error) {
				setError();
			},
			complete: function (resultat, statut) {
				waitHide();
			}
		});
	}
	if (activeTabs() == 1) {
		mask = getSuit();
		$.ajax({
			url: 'api/suitmultipbn',
			type: 'POST',
			data: {
				mask: mask,
				count: count,
				dealer: dealer,
				comment: comment,
				vulnerable: vulnerable
			},
			success: function (result) {
				setSuccess();
				download("donnes.pbn", result)
			},
			error: function (result, statut, error) {
				setError();
			},
			complete: function (resultat, statut) {
				waitHide();
			}
		});
	}

	if (activeTabs() == 2) {
		mask = getPoints();
		$.ajax({
			url: 'api/pointmultipbn',
			type: 'POST',
			data: {
				mask: mask,
				count: count,
				dealer: dealer,
				comment: comment,
				vulnerable: vulnerable
			},
			success: function (result) {
				setSuccess();
				download("donnes.pbn", result)
			},
			error: function (result, statut, error) {
				setError();
			},
			complete: function (resultat, statut) {
				waitHide();
			}
		});
	}


}


function indexAndGenerate() {
	getFromIndex();
}

function razPar() {
	$("#showscore").html("&nbsp;");
	$("#nnt").text("");
	$("#ns").text("");
	$("#nh").text("");
	$("#nd").text("");
	$("#nc").text("");


	$("#ent").text("");
	$("#es").text("");
	$("#eh").text("");
	$("#ed").text("");
	$("#ec").text("");

}

function raz() {
	$("#index-donne").val("");
	razPar();
	$("#okpbn").text("N");
	$('#currentpbn').val("");
	$('#north-spade').text("");
	$('#north-heart').text("");
	$('#north-diamond').text("");
	$('#north-club').text("");

	$('#east-spade').text("");
	$('#east-heart').text("");
	$('#east-diamond').text("");
	$('#east-club').text("");

	$('#south-spade').text("");
	$('#south-heart').text("");
	$('#south-diamond').text("");
	$('#south-club').text("");

	$('#west-spade').text("");
	$('#west-heart').text("");
	$('#west-diamond').text("");
	$('#west-club').text("");
	$('#wp').text("");
	$('#ep').text("");
	$('#sp').text("");
	$('#np').text("");

	if (activeTabs() == 0) {
		$('#n-s').val("");
		$('#n-h').val("");
		$('#n-d').val("");
		$('#n-c').val("");

		$('#e-s').val("");
		$('#e-h').val("");
		$('#e-d').val("");
		$('#e-c').val("");

		$('#s-s').val("");
		$('#s-h').val("");
		$('#s-d').val("");
		$('#s-c').val("");

		$('#w-s').val("");
		$('#w-h').val("");
		$('#w-d').val("");
		$('#w-c').val("");
	}
	if (activeTabs() == 1) {

		$("#c-n-s").val("");
		$("#c-n-h").val("");
		$("#c-n-d").val("");
		$("#c-n-c").val("");

		$("#c-e-s").val("");
		$("#c-e-h").val("");
		$("#c-e-d").val("");
		$("#c-e-c").val("");

		$("#c-s-s").val("");
		$("#c-s-h").val("");
		$("#c-s-d").val("");
		$("#c-s-c").val("");

		$("#c-w-s").val("");
		$("#c-w-h").val("");
		$("#c-w-d").val("");
		$("#c-w-c").val("");
	}
	if (activeTabs() == 2) {

		$("#nmin").val("");
		$("#nmax").val("");
		$("#emax").val("");
		$("#emin").val("");

		$("#smin").val("");
		$("#smax").val("");
		$("#wmin").val("");
		$("#wmax").val("");
	}
}


function setNLToBR(sTab) {
	sTab = sTab.replace(/\n/gi, "<br/>").trim();
	return sTab
}


function setError() {
	$("#okpbn").text("N");
	$('#currentpbn').val("");
	razPar();
	$('#error').show();
	$('#generatebtn').prop("disabled", false);
	$('#getfromindexbtn').prop("disabled", false);
	$('#pbnbtn').prop("disabled", false);
	$('#calcparbtn').prop("disabled", false);
	$('#setpbnbtn').prop("disabled", false);
}

function setSuccess() {
	$("#okpbn").text("Y");
	$('#error').hide();
	$('#generatebtn').prop("disabled", false);
	$('#getfromindexbtn').prop("disabled", false);
	$('#pbnbtn').prop("disabled", false);
	$('#calcparbtn').prop("disabled", false);
	$('#setpbnbtn').prop("disabled", false);
}

function successPost(data) {
	var result = jQuery.parseJSON(data);
	var index = result.index;
	var r = result.pbn;
	$('#np').text("N-Points:" + result.np);
	$('#ep').text("E-Points:" + result.ep);
	$('#sp').text("S-Points:" + result.sp);
	$('#wp').text("O-Points:" + result.wp);
	$("#currentpbn").val(r);
	$('#index-donne').val(index);
	setSuccess();
	showDesk(r);
	$('#generatebtn').prop("disabled", false);
	$('#getfromindexbtn').prop("disabled", false);
	$('#pbnbtn').prop("disabled", false);
	$('#calcparbtn').prop("disabled", false);
	$('#setpbnbtn').prop("disabled", false);
}

function rot() {
	if ($("#okpbn").text() == "Y") {
		var mask
		mask = $("#currentpbn").val();
		$("#currentpbn").val("E:" + mask);
		setPbn();
	}
}

function initCalc() {
	waitShow();
	$('#error').hide();
	razPar();
}

function setPbn() {
	initCalc();
	var mask
	mask = $("#currentpbn").val();
	$('#setpbnbtn').prop("disabled", true);
	$.ajax({
		url: 'api/maskpbn',
		type: 'POST',
		data: {
			mask: mask
		},
		success: function (data) {
			if (data != "") {
				successPost(data);
			}
		},
		error: function (result, statut, error) {
			setError();
		},
		complete: function (resultat, statut) {
			waitHide();
		}
	});

}


function getIdemList(result){
	s1=result.ewl
	s1 = s1.replace(/EW/gi, "NS").trim();
	if (s1=result.nsl){
		return ""
	}
	return s1
}

function calcPar() {
	if ($("#okpbn").text() == "Y") {
		initCalc();
		var mask = "N:" + $("#currentpbn").val();
		var vul = $("#vulnerable").find('option:selected').val()
		$('#calcparbtn').prop("disabled", true);
		$.ajax({
			url: 'api/parpbn',
			type: 'GET',
			data: {
				pbn: mask,
				vul: vul
			},
			success: function (data) {
				if (data != "") {
					var result = jQuery.parseJSON(data);
					var par = "Points " + result.nss + " " + result.ews.replace(/EW/gi, "EO").trim() + " Contrats " + translateScore(result.nsl) + " "+translateScore(getIdemList(result)) ;
					$("#showscore").text(par);
					$("#nnt").text(result.nnt);
					$("#ns").text(result.ns);
					$("#nh").text(result.nh);
					$("#nd").text(result.nd);
					$("#nc").text(result.nc);

					$("#ent").text(result.ent);
					$("#es").text(result.es);
					$("#eh").text(result.eh);
					$("#ed").text(result.ed);
					$("#ec").text(result.ec);

				}
			},
			error: function (result, statut, error) {
				setError();
			},
			complete: function (resultat, statut) {
				waitHide();
				$('#calcparbtn').prop("disabled", false);
			}
		});
	}
}

function generate() {
	initCalc();
	$('#generatebtn').prop("disabled", true);
	if (activeTabs() == 0) {
		var mask
		mask = getPbn();
		$.ajax({
			url: 'api/maskpbn',
			type: 'POST',
			data: {
				mask: mask
			},
			success: function (data) {
				if (data != "") {
					successPost(data)
				}
			},
			error: function (result, statut, error) {
				setError();
			},
			complete: function (resultat, statut) {
				waitHide();
			}
		});
	}
	if (activeTabs() == 1) {
		var mask
		mask = getSuit();
		$.ajax({
			url: 'api/suitpbn',
			type: 'POST',
			data: {
				mask: mask
			},
			success: function (data) {
				if (data != "") {
					successPost(data)
				}
			},
			error: function (result, statut, error) {
				setError();
			},
			complete: function (resultat, statut) {
				waitHide();
			}
		});
	}
	if (activeTabs() == 2) {
		var mask
		mask = getPoints();
		$.ajax({
			url: 'api/pointpbn',
			type: 'POST',
			data: {
				mask: mask
			},
			success: function (data) {
				if (data != "") {
					successPost(data)
				}
			},
			error: function (result, statut, error) {
				setError();
			},
			complete: function (resultat, statut) {
				waitHide();
			}
		});
	}
}

function activeTabs() {
	var $tabs = $('#tabs').tabs();
	var active = $tabs.tabs('option', 'active');
	return active;
}

function retValue(t) {
	var r = 0;
	var v = $(t).val()
	if (jQuery.isNumeric(v)) {
		r = v;
	}
	r = parseInt(r, 10);
	return r
}

function setNotVisible(v, a) {
	if (a) {
		$(v).hide();
	} else {
		$(v).show();
	}
}

function viewHand(hand) {
	if (hand == "N") {
		setNotVisible("#north-spade", ($("#h-n").is(':checked')));
		setNotVisible("#north-heart", ($("#h-n").is(':checked')));
		setNotVisible("#north-diamond", ($("#h-n").is(':checked')));
		setNotVisible("#north-club", ($("#h-n").is(':checked')));
		setNotVisible("#np", ($("#h-n").is(':checked')));
	}
	if (hand == "E") {
		setNotVisible("#east-spade", ($("#h-e").is(':checked')));
		setNotVisible("#east-heart", ($("#h-e").is(':checked')));
		setNotVisible("#east-diamond", ($("#h-e").is(':checked')));
		setNotVisible("#east-club", ($("#h-e").is(':checked')));
		setNotVisible("#ep", ($("#h-e").is(':checked')));
	}
	if (hand == "S") {
		setNotVisible("#south-spade", ($("#h-s").is(':checked')));
		setNotVisible("#south-heart", ($("#h-s").is(':checked')));
		setNotVisible("#south-diamond", ($("#h-s").is(':checked')));
		setNotVisible("#south-club", ($("#h-s").is(':checked')));
		setNotVisible("#sp", ($("#h-s").is(':checked')));
	}
	if (hand == "W") {
		setNotVisible("#west-spade", ($("#h-w").is(':checked')));
		setNotVisible("#west-heart", ($("#h-w").is(':checked')));
		setNotVisible("#west-diamond", ($("#h-w").is(':checked')));
		setNotVisible("#west-club", ($("#h-w").is(':checked')));
		setNotVisible("#wp", ($("#h-w").is(':checked')));
	}

}

$(document).ready(function () {
	$("#tabs").tabs();
	getVersion();
});