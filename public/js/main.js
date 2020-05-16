
var FromDate;
var ToDate;

var PassDate;

const visitorsTbBody=document.querySelector("#visitorsTb>tbody");
console.log(visitorsTbBody)
const tableVisitors = $('#visitorsTb');


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

//var formUser = document.querySelector('#formValue').value;



/* Daily Users*/
//document.querySelector('#btn1').onclick = makeCharts;

function makeCharts() {
    FromDate = document.getElementById("fromDate").value;
    ToDate = document.getElementById("toDate").value;
    PassDate = "fromDate" + FromDate + "&toDate" + ToDate;

    alert(FromDate);

    alert(FromDate);

    ToDate


    console.log(PassDate);
    alert(PassDate);
    var url = '/dailyUsers'

    const xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-url-urlencoded")
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            try{
                const dataJson = JSON.parse(xhr.responseText);
                console.log(dataJson);
                visitorsTable(dataJson);

                const labels = dataJson.map(function (e) {
                    return e.captured_time;
                });
                const data = dataJson.map(function (e) {
                    return e.num_of_user;

                });
                const ctx = document.querySelector("#myChart").getContext('2d');
                new Chart(ctx, {
                    type: 'bar',
                    data: {
                        labels: labels,
                        datasets: [{
                            label: 'DailyUsers',
                            backgroundColor: 'rgb(54, 104, 115)',
                            borderColor: 'rgb(54, 104, 115)',
                            data: data,
                        }]
                    }
                });

            }catch (e) {
                console.warn("could not load data ");

            }






        }





    };





    xhr.send(PassDate);

}
function visitorsTable(json) {
    console.log(json);

    //clear out existing table data
    while (tableVisitorsBody.firstChild) {
        tableVisitorsBody.removeChild(tableVisitorsBody.firstChild);
    }
//populate table
    // let matrix=[];
    json.forEach((test) => {
        tableVisitors.append("<tr>" +
            "<td class='col-xs-3'>" + "<h6>" + test.captured_time + "</h6></td> " +
            "<td class='col-xs-2'> <h6>" + test.num_of_user + "</h6>" + "</td>" +

            "</tr>");  });
    const csvData =generateCsv(json);
    download(csvData);
}

const generateCsv=function (jsonCsv) {
    const csvRows =[];

    //get the Headers
    const headers=Object.keys(jsonCsv[0]);
    csvRows.push(headers.join(','));


    //loop over the rows
    for(const row of jsonCsv){

        const values = headers.map(header =>{
            const escaped=(''+row[header]).replace(/"/g,'\\"');
            return `"${escaped}"`;
        });
       csvRows.push(values.join(','));
    }
    return csvRows.join('\n');

    //from escaped comma separated values

};
const download = function (downloadData) {
    const blob = new Blob([downloadData],{type:'text/csv'});
    const url= window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.setAttribute('hidden','');
    a.setAttribute('href',url);
    a.setAttribute('download','visitors.csv');
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);

}

// function generateCsv(jsonCsv){
//     const data =jsonCsv.map(row => ({
//         dateCsv:row.captured_time,
//         visitorsCsv:row.num_of_user,
//
//     }));
//     const csvData=objectToCsv
//
// }
    //console.log(matrix);



/* The function */

/*avgDwelTime*/
/*
$(dataJson).each(function (i,visitorsTb) {
    $(#tableVisitorsBody).append($("<tr>")
        .append($("<td>").append(visitorsTb.captured_time))
        .append($("<td>").append(visitorsTb.num_of_user))
    )

});
//document.querySelector('#btn').onclick = makeCharts2;
/*function makeCharts2() {
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
/*function makeCharts3() {
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
*/
/*

$(document).ready(function () {
    $('#btn1').click(function () {
        // var vars={
        //     FromDate:$('#fromDate').val(),
        //     ToDate :$('#toDate').val()
        // }
      // console.log(vars);
        var url="/avgDTime";
        $.get(url,function (data,status,xhr) {
            console.log(data);
            console.log(status);
            console.log(xhr);
         // var jsonData=JSON.parse(data);
            $.each(data,function (i,itme) {
                var lables = data[i].captured_time;
                var data1=data[i].num_of_user;

            });



            var ctx = document.querySelector("#myChart").getContext('2d');
            new Chart(ctx, {
                type: 'bar',
                data: {
                    labels: labels,
                    datasets: [{
                        label: 'DailyUsers',
                        backgroundColor: 'rgb(54, 104, 115)',
                        borderColor: 'rgb(54, 104, 115)',
                        data: data1,
                    }]
                }
            });

        })


    });

    })
*/


//test
