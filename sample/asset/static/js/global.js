
$(document).ready(function() {
    $.fn.select2.defaults.set("theme", "bootstrap");        
    initAll();
    $('nav').on('click', function() {
        elems = $('nav ul li');
        if (elems[0].style.bottom === "") {
            $('nav span.entry').html('<i class="fa fa-times">');            
            t = 0;
            for (i = 0 ; i < elems.length ; i++) {
                t += 70;
                elems[i].style.bottom = t + "px";
            }
        } else {
            $('nav span.entry').html('<i class="fa fa-bars">'); 
            for (i = 0 ; i < elems.length ; i++) {
                elems[i].style.bottom = "";
            }
        }
    });
});