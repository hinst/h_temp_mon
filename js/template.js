var readTemperature = function() {
	jQuery.ajax({
		type: "GET",
		url: AppURL + "/current_temperature",
		dataType: "xml",
		success: function(xmlDocument) {
			console.log(xmlDocument != null);
			var temperature = $(xmlDocument).find("Temperature").text();
			console.log("temperature=" + temperature);
		}
	});
}
readTemperature();