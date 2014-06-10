var ajax  = {
    diff : 0,
    path : '/api/web',
    timeout : 30000,
    errorHandler : null,

    setServerTime : function(unix){
        this.diff = unix>0 ? (new Date()).getTime() - unix*1000 : 0;
    },

    serverTime : function(time){
        return (new Date()).getTime() - this.diff;
    },

    getHostName : function(){
        var hostArr = window.location.host.split('.')
        return hostArr[hostArr.length-2]+'.'+hostArr[hostArr.length-1]
    },

    getUrlVars: function(){
        var vars = [], hash;
        var hashes = window.location.href.slice(window.location.href.indexOf('?') + 1).split('&');
        for(var i = 0; i < hashes.length; i++) {
        hash = hashes[i].split('=');
        vars.push(hash[0]);
        vars[hash[0]] = hash[1];
        }
        return vars;
    },

    doBusy : function(el, sw){
        var container,overlay
        
        if(typeof el == "string")
            container  = $(el);
        else
            container  = el;

        var options = {
            bgColor         : '#333',
            duration        : 200,
            opacity         : 0.7
        }
        container.each(function(){
            if(sw){
                overlay = $('<div></div>').css({
                        'background-color': options.bgColor,
                        'opacity':options.opacity,
                        'width':$(this).width(),
                        'height':$(this).height(),
                        'position':'absolute',
                        'top':'0px',
                        'left':'0px',
                        'z-index':9999
                })
                overlay = $('<div class="block-overlay"></div>').css({
                        'position': 'relative'
                }).append(overlay)

                $(this).prepend(
                    overlay.append('<div class="bloack-ui"></div>').fadeIn(options.duration)
                );
            }else{
                overlay = $(this).children(".block-overlay");
                if (overlay.length>0) {
                    // overlay.fadeOut(options.duration, function() {
                    //     overlay.remove();
                    // });
                    overlay.remove();
                }
            }
        })
    },

    sendAlone : function(method, args, callback, blockTo){
        if(this.deferred && this.deferred.state()=='pending'){
            var func = function(){}
            return {done:func, fail:func, always:func};
        }
        this.deferred = this.send(method, args, callback, blockTo);
        return this.deferred;
    },

    send : function(method, args, callback, blockTo){
        var pm = {method:method, args:args===undefined?null:args};
        var options = {type:'POST', dataType:'json',cache:false, timeout: this.timeout},k = null

        var busyFunc = this.doBusy
        if (blockTo)
            busyFunc(blockTo, true);
        var clas = this;
            deferred = $.ajax({
                url:        this.path,
                type:       options['type'],
                dataType:   options['dataType'],
                cache:      options['cache'],
                //jsonp: 'callback',
                data:       {json: JSON.stringify(pm)},
                timeout:    options['timeout']
            });

        deferred.always(function(){
            if(blockTo) 
                busyFunc(blockTo, false);

        }).done(function(data,statusText, jqXHR ){
            if(callback)
                callback.call(clas, data, args)
        }).fail(function (jqXHR, textStatus, errorThrown){
            switch(jqXHR.status){
                case 599:
                    if(clas.errorHandler){
                        try{
                            clas.errorHandler(JSON.parse(jqXHR.responseText))
                        }catch(e){
                            console.log(e)
                        }
                    }
                    break;
                case 404:
                    alert(textStatus+': api not found!')
                    break;
                default:
                    alert(textStatus+': '+jqXHR.responseText)
                    break;
            }
        });

        return deferred;
    },

    datasetDecode : function(data){
        if(!data || data.length==0)
            return data

        var fields=data[0], l = data.length, obj, value, items=[]
        for(var i=1; i < l; i++) {
            obj = new Object()
            for(var k=0; k < fields.length; k++)
                obj[fields[k]] = data[i][k]
            items.push(obj)
        }
        return items;
    },

    setCookie : function (name,value,days) {
        if (days) {
            var date = new Date();
            date.setTime(date.getTime()+(days*86400000));
            var expires = "; expires="+date.toGMTString();
        }
        else var expires = "";
        document.cookie = name+"="+value+expires+"; path=/";
    },
    getCookie: function (name) {
        var nameEQ = name + "=";
        var ca = document.cookie.split(';');
        for(var i=0;i < ca.length;i++) {
            var c = ca[i];
            while (c.charAt(0)==' ') c = c.substring(1,c.length);
            if (c.indexOf(nameEQ) == 0) return c.substring(nameEQ.length,c.length);
        }
        return null;
    }

};
