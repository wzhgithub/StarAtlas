window.onload = function() {
	var dl = document.getElementsByTagName('dl')[0];
	var dds = dl.getElementsByTagName('dd');
	jiazai1(dds, dds.length - 1);
	jiazai(dds, dds.length - 1);
	dl.style.transform = "rotateX(-10deg) rotateY(0deg)";
	window.onmousedown = function(e) {
		e = e || window.event;
		var mouseX = e.clientX,
			mouseY = e.clientY;
		var xx = dl.style.transform.substring(dl.style.transform.indexOf('rotateX(') + 8)
		xx = parseInt(xx.substring(0, xx.indexOf('deg')));
		var yy = dl.style.transform.substring(dl.style.transform.indexOf('rotateY(') + 8)
		yy = parseInt(yy.substring(0, yy.indexOf('deg')));
		window.onmousemove = function(e) {
			e = e || window.event;
			var x1 = e.clientX - mouseX,
				y1 = e.clientY - mouseY;
			var Xdeg = xx - y1 / 6;
			var Ydeg = yy + x1 / 6;
			if(Xdeg > 360 || Xdeg < -360) {
				Xdeg %= 360;
			}
			if(Ydeg > 360 || Ydeg < -360) {
				Ydeg %= 360;
			}
			dl.style.transform = "rotateX(" + Xdeg + "deg) rotateY(" + Ydeg + "deg)";
		}
		window.onmouseup = function() {
			window.onmousemove = null;
		}
	}
}

function jiazai1(dds, n) {
	var img = dds[n].getElementsByTagName('img')[0];
	var div = document.createElement('div');
	div.style.background = "-webkit-linear-gradient(rgb(0, 0, 0) 50%, rgba(255, 255, 255, 0)),url(" + img.src + ")";
	div.style.backgroundSize = "100% 100%";
	div.style.backgroundPosition = "center center";
	div.className = 'yy';
	dds[n].appendChild(div);
	if(n > 0) {
		jiazai1(dds, n - 1);
	}
}

function jiazai(dds, n) {
	var speed = 100;
	var translateZTerminus = 400;
	var angle = 360 / dds.length * n;
	var translateZ = 0;
	var rotateY = 0;
	var time = setInterval(function() {
		translateZ += translateZTerminus / speed * 10;
		rotateY += angle / speed * 10;
		if(rotateY >= angle && translateZ >= translateZTerminus) {
			clearInterval(time);
			dds[n].style.transform = 'rotateY(' + angle + 'deg) translateZ(' + translateZTerminus + 'px)';
			if(n > 0) {
				jiazai(dds, n - 1);
			}
		}
		dds[n].style.transform = 'rotateY(' + rotateY + 'deg) translateZ(' + translateZ + 'px)';
	}, 10);
}