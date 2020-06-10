$(function () {
    getCity();
})

function getCity() {
    $.ajax({
        url: 'https://free-api.heweather.net/s6/weather/now?location=58.240.228.66&key=db86a5196f304e52a4369818c5182e60',
        type: 'POST',
        dataType: 'json',
        success:function(data) {
            $('.weather').css("background","url(https://cdn.heweather.com/cond_icon/"+data.HeWeather6[0].now.cond_code+".png)")
        }
    });
}