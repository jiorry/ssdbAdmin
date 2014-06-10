var BaseControl = function(ctype){
	var thisClas = this
	this.ctype = ctype;
	this.name = '';

	this.init = function(){
		if(this.ctype=='hsetItem'||this.ctype=='zsetItem')
			this.name = $('#txtName').val()

		$('#btnQuery')
			.click(function(){
				if(thisClas.ctype == 'zsetItem'){
					thisClas.loadZsetItems($('#txtKeyStart').val(),$('#txtStart').val(), $('#txtEnd').val(), $('#txtLimitCount').val())
					return;
				}
				thisClas.loadRecord($('#txtStart').val(), $('#txtEnd').val(), $('#txtLimitCount').val())
			})
			.trigger('click');

		var $table = $('#resultTable');

		$('tbody', $table).on('click', 'a[data-action=view]', function(){
			thisClas.viewValue($(this).parent().parent().data('key'))
		})

		$('tbody', $table).on('click', 'a[data-action=delete]', function(){
			var key = $(this).parent().parent().data('key');

			if(confirm('do you want to delete record [' + key + ']?'))
				thisClas.del(key);
		})

		$table.addClass('table-'+ctype);
	}

	this.loadZsetItems = function(keyStart, start, end, limit){
		ajax.send('LoadZsetItems', {name: this.name, keyStart:keyStart, start:start, end:end, limit:parseInt(limit)})
		.done(function(result){
			var i,html = ''
			for (i = 0; i < result.length; i++) {
				html += thisClas.itemHtml(result[i]);
			};
			
			$('#resultTable tbody').html(html)
		})
	}
	this.loadRecord = function(start, end, limit){
		start = start || '';
		end = end || '';
		limit = limit || 500;

		ajax.send('LoadRecord', {name: this.name, ctype:ctype, start:start, end:end, limit:parseInt(limit)})
		.done(function(result){
			var i,html = ''
			for (i = 0; i < result.length; i++) {
				html += thisClas.itemHtml(result[i]);
			};
			
			$('#resultTable tbody').html(html)
		})
	}

	this.itemHtml = function(item){
		var h;
		switch(this.ctype){
			case 'hset':
				h = '<a target="_blank" href="/hset/get?q='+encodeURIComponent(item)+'"><i class="glyphicon glyphicon-eye-open"></i> view</a>';
				break;
			case 'zset':
				h = '<a target="_blank" href="/zset/get?q='+encodeURIComponent(item)+'"><i class="glyphicon glyphicon-eye-open"></i> view</a>';
				break;
			default:
				h = '<a data-action="view"><i class="glyphicon glyphicon-eye-open"></i> view</a>';
				break;
		}
		return '<tr data-key="'+item+'"><td>'+item+'</td>'+
			'<td>'+h+'</td>'+
			'<td class="buttons"><a data-action="delete" class="glyphicon glyphicon-remove"></a> <a data-action="edit" class="glyphicon glyphicon-edit"></a></td></tr>'
	}

	this.viewValue = function(key){
		ajax.send('GetValue', {ctype:ctype, key:key, name:this.name})
			.done(function(result){
				$('#resultTable tbody tr[data-key='+key+']').find('td:nth-child(2)').text(result)
			})
			
	}

	this.del = function(key){
		ajax.send('Del', {ctype:ctype, key:key, name: this.name})
			.done(function(result){
				$('#resultTable tbody tr[data-key='+key+']').remove()
			})
	}
}

var keys = new BaseControl("key");
var hset = new BaseControl("hset");
var zset = new BaseControl("zset");
	
var hsetItem = new BaseControl("hsetItem");
var zsetItem = new BaseControl("zsetItem");
