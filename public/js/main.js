const btn1 = $('#btn1');
btn1.on('click', () => {
    const chartRow1 = $('#chartRow1');
    const row1 = $('#row1');
    row1.toggleClass('border border-bottom-0 border-primary');
    chartRow1.toggleClass('d-none');
    btn1.toggleClass('btn-success');
    if (chartRow1.hasClass('d-none')) {
        btn1.html('More Details');
    } else {
        btn1.html('Show Less');
    }
});
/* Daily Users*/
//document.querySelector('#btn1').onclick = makeCharts;
function makeCharts() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET','/dailyUsers', true);
    xhr.onreadystatechange = function() {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200){
            var dataJson = JSON.parse(xhr.responseText)

            var labels=dataJson.map(function (e) {
                return e.captured_time;

            });
            var data =dataJson.map(function (e) {
                return e.num_of_user;

            });
            var ctx = document.querySelector("#myChart").getContext('2d');
            new Chart(ctx, {
                type: 'bar',
                data: {
                    labels:labels,
                    datasets: [{
                        label: 'DailyUsers',
                        backgroundColor: 'rgb(54, 104, 115)',
                        borderColor: 'rgb(54, 104, 115)',
                        data:data,
                    }]            }
            });
        }
    };
    xhr.send();

}

/*avgDwelTime*/
//document.querySelector('#btn').onclick = makeCharts2;
function makeCharts2() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET','/avgDTime', true);
    xhr.onreadystatechange = function() {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200){
            var dataJson = JSON.parse(xhr.responseText)

            var labels=dataJson.map(function (e) {
                return e.cpatured_date;

            });
            var data =dataJson.map(function (e) {
                return e.captured_time_dw;

            });
            var ctx = document.querySelector("#myChartDailyUsers").getContext('2d');
            new Chart(ctx, {
                type: 'line',
                data: {
                    labels:labels,
                    datasets: [{
                        label:'DwTime',
                       // backgroundColor:'rgb( 99, 154, 166)',
                        borderColor: 'rgb(6,14,76)',
                        fill:false,


                        data:data,
                    }]
                }
            });
        }
    };
    xhr.send();

}
/*rpUsers*/
//document.querySelector('h1').onclick = makeCharts3;
function makeCharts3() {
    var xhr = new XMLHttpRequest();
    xhr.open('GET','/rpUsers', true);
    xhr.onreadystatechange = function() {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200){
            var dataJson = JSON.parse(xhr.responseText)

            var labels=dataJson.map(function (e) {
                return e.p_date;

            });
            var data =dataJson.map(function (e) {
                return e.rp_count;

            });
            var ctx = document.querySelector("#myChart3").getContext('2d');
            new Chart(ctx, {
                type: 'bar',
                data: {
                    labels:labels,
                    datasets: [{
                        label: 'RpUsers',
                        backgroundColor: 'rgb( 25, 60, 64)',
                        borderColor: 'rgb( 25, 60, 64)',
                        data:data,
                    }]            }
            });
        }
    };
    xhr.send();

}

