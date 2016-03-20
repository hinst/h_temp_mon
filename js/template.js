var lastTemperature = "";
var degreeString = "°";
var readTemperature = function() {
	jQuery.ajax({
		type: "GET",
		url: AppURL + "/current_temperature",
		dataType: "xml",
		success: function(xmlDocument) {
			var temperature = $(xmlDocument).find("Temperature").text();
			lastTemperature = temperature;
			$('#temperature').html("" + temperature + degreeString);
		},
		error: function (a, b, c) {
			$('#temperature').html("<s>" + lastTemperature + degreeString + "</s>");
		}
	});
}
var readTemperatureContinuously = function() {
	readTemperature();
	setTimeout(readTemperatureContinuously, 2000);
}
readTemperatureContinuously();