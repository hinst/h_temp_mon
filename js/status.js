var readTemperature = function() {
	jQuery.ajax({
		type: "GET",
		url: "{{.AppURL}}/current_temperature",
		dataType: "xml",
		success: function(xml) {
			var xmlDocument = jQuery.parseXML(xml);
			var temperature = $(xmlDocument).find("Temperature");
			console.log("temperature=" + temperature);
		}
	});
}
readTemperature();