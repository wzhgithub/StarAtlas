$(function() {
	$('dl').transform('rotateX', '10deg')
	$('dl').transform('rotateY', '0deg')
	jiazai1($('dd'), $('dd').length - 1);
	jiazai($('dd'), $('dd').length - 1);
	var rotate = {
		X: 0,
		Y: 0
	};
	var t = 0;
	var inertiaTime = null;
	$(window).mousedown(function(e) {
		clearInterval(inertiaTime);
		$('dl').transform('rotateX', parseFloat($('dl').transform('rotateX')) % 360 + 'deg')
		$('dl').transform('rotateY', parseFloat($('dl').transform('rotateY')) % 360 + 'deg')
		var mouseX = e.clientX,
			mouseY = e.clientY;
		t = new Date().getTime();
		rotate.X = parseFloat($('dl').transform('rotateX'))
		rotate.Y = parseFloat($('dl').transform('rotateY'))
		$(window).mousemove(function(e) {
			var x = rotate.X - (e.clientY - mouseY) / 6;
			var y = rotate.Y + (e.clientX - mouseX) / 6;
			$('dl').transform('rotateX', x + 'deg')
			$('dl').transform('rotateY', y + 'deg')
			//			rotate.X = x;
			//			rotate.Y = y;
			//			mouseX = e.clientX;
			//			mouseY = e.clientY;
			//			t = new Date().getTime();
			//						
		});
	}).mouseup(function() {
		$(window).off('mousemove');
		rotate.X = parseFloat($('dl').transform('rotateX')) - rotate.X;
		rotate.Y = parseFloat($('dl').transform('rotateY')) - rotate.Y;
		t = new Date().getTime() - t;
		t = t < 1 ? 1 : t;
		var juli = {
			X: rotate.X / t * 20,
			Y: rotate.Y / t * 20
		};
		clearInterval(inertiaTime);
		if(Math.abs(juli.X) > 3 || Math.abs(juli.Y) > 3) {
			inertiaTime = setInterval(function() {console.log(juli)
				$('dl').transform('rotateX', juli.X + parseFloat($('dl').transform('rotateX')) + 'deg');
				$('dl').transform('rotateY', juli.Y + parseFloat($('dl').transform('rotateY')) + 'deg');
				juli.X *= 0.9;
				juli.Y *= 0.9;
				if(Math.abs(juli.X) < 1 && Math.abs(juli.Y) < 1) {
					clearInterval(inertiaTime);
					$('dl').transform('rotateX', parseFloat($('dl').transform('rotateX')) % 360 + 'deg')
					$('dl').transform('rotateY', parseFloat($('dl').transform('rotateY')) % 360 + 'deg')
				}
			}, 100);
		}
		$('dl').transform('rotateX', parseFloat($('dl').transform('rotateX')) % 360 + 'deg')
		$('dl').transform('rotateY', parseFloat($('dl').transform('rotateY')) % 360 + 'deg')
	});
});

function jiazai1(dds, n) {
	$('<div></div>', {
		'class': 'yy'
	}).css({
		'background': "-webkit-linear-gradient(rgb(0, 0, 0) 50%, rgba(255, 255, 255, 0)),url(" + dds.eq(n).children('img').attr('src') + ")",
		'background-size': '100% 100%',
		'background-position': "center center"
	}).appendTo(dds.eq(n));
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
			dds.eq(n).transform('rotateY', angle + 'deg')
			dds.eq(n).transform('translateZ', translateZTerminus + 'px');
			if(n > 0) {
				jiazai(dds, n - 1);
			}
		}
		dds.eq(n).transform('rotateY', rotateY + 'deg')
		dds.eq(n).transform('translateZ', translateZ + 'px')
	}, 10);
}

$.fn.transform = function(attr, value) {
	var attrName = attr + '(';
	var transform = $(this)[0].style.transform;
	var attrValue = transform.substr(transform.indexOf(attrName) + attrName.length);
	attrValue = attrValue.substring(0, attrValue.indexOf(')'));
	if(value || value === 0) {
		var valueStr = attr + "(" + value + ")";
		if(transform.indexOf(attr) > -1) {
			var str = transform.substr(transform.indexOf(attr))
			str = str.substring(0, str.indexOf(')') + 1);
			$(this)[0].style.transform = transform.replace(str, valueStr);
		} else {
			if($(this).css('transform') == 'none')
				$(this)[0].style.transform = valueStr
			else
				$(this)[0].style.transform = $(this)[0].style.transform + valueStr
		}
	}
	return attrValue;
}