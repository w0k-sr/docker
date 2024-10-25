window.addEventListener('DOMContentLoaded', function(){

  // テキストエリアのHTML要素を取得
  let textarea_contact = document.getElementById("srow_0");

  // イベント「change」を登録
  textarea_contact.addEventListener("change",function(){
    console.log("Change action");
    console.log(this.value);
  });
});