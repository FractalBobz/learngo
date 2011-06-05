$(function() {
	$('#run').click(function() {
		$('#output').text('Running...');
		$.ajax('/test', {
			contentType: 'text/plain',
			data: $('#code').val(),
			dataType: 'text',
			success: function(data) {
				$('#output').text(data);
			},
			error: function() {
				$('#output').text('Error running.');
			},
			type: 'POST',
		});
	});
});
