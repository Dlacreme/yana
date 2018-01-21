function productcreateaddProduct() {
    $('#create-product-composition-template').clone().removeClass("hidden").removeAttr("id").appendTo("#create-product-composition");
    window.setTimeout(function() {
        setName();
        initSelect2();
    });
}
function productcreateremove(e) {
    var p = e.target.parentNode;
    while (p.tagName != "DIV") {
        p = p.parentNode;
    }
    p.remove();
    setName();
}
function setName() {
    var eselect = $('.var-select-item');
    var eq = $('.var-quantity-item');

    for (i = 0 ; i < eselect.length ; i++) {
        if (i == 0) {
            continue;
        }
        $(eselect[i]).attr('name', (i-1).toString()+".stock");
    }
    for (i = 0 ; i < eq.length ; i++) {
        if (i == 0) {
            continue;
        }
        $(eq[i]).attr('name', (i-1).toString() + ".quantity");
    }
}
function addCompoToProduct() {
    var products = $('#input-compo');
    var quantity = $('#input-quantity');

    //$('#template-form').append('<div class="col-md-12 item-compo"><div class="col-md-9">'+products.select2().text()+' - ' + quantity.val()  +' </div><div class="col-md-2"><button type="button" class="btn btn-danger" onclick="Dynamo.Html.RemoveSelf(event, \'item-compo\')"><i class="fa fa-trash"></i></div></div><div class="col-md-1 hidden"><input class="hidden" type="text" value="{'+products.select2().val()+', '+quantity.val()+'}" /></div></div>');
    $('#template-form').append('<div class="col-md-12 item-compo"><div class="col-md-9">'+products.select2().text()+' - ' + quantity.val()  +' </div><div class="col-md-1 hidden"><input class="hidden" type="text" value="{'+products.select2().val()+', '+quantity.val()+'}" /></div></div>');
}