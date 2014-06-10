var BaseControl = function(ctype){
	var thisClas = this
	this.ctype = ctype;

	this.init = function(){
		this.loadRecord();
		$('#btnQuery').click(function(){
			thisClas.loadRecord($('#txtStart').val(), $('#txtEnd').val(), $('#txtLimitCount').val())
		})

		var $table = $('#resultTable');

		$('tbody', $table).on('click', 'a[data-action=view]', function(){
			thisClas.viewValue($(this).parent().parent().data('name'))
		})

		$('tbody', $table).on('click', 'a[data-action=delete]', function(){
			var name = $(this).parent().parent().data('name');

			if(confirm('do you want to delete record [' + name + ']?'))
				thisClas.del(name);
		})

		$table.addClass('table-'+ctype);
	}

	this.loadRecord = function(start, end, limit){
		start = start || '';
		end = end || '';
		limit = limit || 500;

		ajax.send('LoadRecord', {ctype:ctype, start:start, end:end, limit:parseInt(limit)})
		.done(function(result){
			var i,html = ''
			for (i = 0; i < result.length; i++) {
				html += '<tr data-name="'+result[i]+'"><td>'+result[i]+'</td><td><a data-action="view"><i class="glyphicon glyphicon-eye-open"></i> view</a></td><td class="buttons"><a data-action="delete" class="glyphicon glyphicon-remove"></a> <a data-action="edit" class="glyphicon glyphicon-edit"></a></td></tr>'
			};
			
			$('#resultTable tbody').html(html)
		})
	}

	this.viewValue = function(name){
		if(this.ctype=='key'){
			ajax.send('GetValue', {ctype:ctype, name:name})
				.done(function(result){
					$('#resultTable tbody tr[data-name='+name+']').find('td:nth-child(2)').text(result)
				})
		}
			
	}

	this.del = function(name){
		ajax.send('Del', {ctype:ctype, name:name})
			.done(function(result){
				$('#resultTable tbody tr[data-name='+name+']').remove()
			})
	}
}

var keys = new BaseControl("key");
var hset = new BaseControl("hset");
var zset = new BaseControl("zset");
	
