$(document).on("change", ".sudoku-child", function() {
  console.log( $(this).val() );
  console.log( $(this).data("col") );
  console.log( $(this).data("row") );
  updateAnswer($(this).data("col"), $(this).data("row"), $(this).val())
});
// 入力データの更新
function updateAnswer(col_val, row_val, num_val) {
  $.ajax( {
    url: '/sudoku/update',
    type: "POST",
    dataType: 'json',
    data: {
      id : Number($('input[id="sudoku_id"]').val()),
      col: Number(col_val),
      row: Number(row_val),
      num: Number(num_val)
    },
    success: function(response){
      // 間違ったら、背景は赤
      $('[data-col="' + col_val + '"][data-row="' + row_val + '"]')
      .attr('data-check', response.data.flag)
      // コメントと猫を出しわける
      $('.neko').eq($('input[id="comment_num"]').val()-1).attr('data-neko', true)
      $('.neko').eq(response.data.comment - 1).attr('data-neko', false)
      $('input[id="comment_num"]').val(response.data.comment)
    },
    error: function(xhr, textStatus, errorThrown){
      console.log( xhr );
      console.log( textStatus );
      console.log( errorThrown );
      alert('Error');
    }
  });
}
