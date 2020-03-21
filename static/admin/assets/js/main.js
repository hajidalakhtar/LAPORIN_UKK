// $.ajax({
//     url: "demo_test.txt",
//     success: function (result) {
//         $("#div1").html(result);
//     }
// });

function gantiStatus(sel) {
    sel.submit();
    // setTimeout(() => {
    //     $('.submit_ganti_status').click();

    // }, 100);
}


function OpenModalEditUser() {

    let FullName = event.target.getAttribute('data-FullName');
    let Role = event.target.getAttribute('data-Role');
    let Bidang = event.target.getAttribute('data-Bidang');
    let Gender = event.target.getAttribute('data-Gender');
    let NoTelp = event.target.getAttribute('data-NoTelp');
    let alamat = event.target.getAttribute('data-Alamat');
    let Foto = event.target.getAttribute('data-Foto');
    let Email = event.target.getAttribute('data-Email');
    let username = event.target.getAttribute('data-Username');

    console.log("email"+Email)

    $("#FullName").val(FullName)
    $("#Role").val(Role)
    $("#Bidang").val(Bidang)
    $("#Gender").val(Gender)
    $("#NoTelp").val(NoTelp)
    $("#Alamat").val(alamat)
    $("#Foto").val(Foto)
    $("#Email").val(Email)
    $("#Username").val(username)


    // $('#edit_modal').modal('show')
}
