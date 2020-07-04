
var FromDate;
var ToDate;

var PassDate;

const visitorsTbBody=document.querySelector("#visitorsTb>tbody");
console.log(visitorsTbBody)
const tableVisitors = $('#visitorsTb');
const tableDwVisitors = $('#DwVisitorsTb');
const tableRpVisitors = $('#rpVisitorsTb');
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
/*Date range picker jQuery*/
function passDates(datefrom1,dateTo1) {
    //console.log('New date range selected: ' +datefrom1 + ' to ' + dateTo1 );
    PassDate = "fromDate" + datefrom1 + "&toDate" + dateTo1;
    return PassDate;


}


function makeCharts() {
    // FromDate = document.getElementById("fromDate").value;
    // ToDate = document.getElementById("toDate").value;
    // PassDate = "fromDate" + FromDate + "&toDate" + ToDate;
    console.log(PassDate)

    console.log(PassDate);

    var url = '/dailyUsers'

    const xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-url-urlencoded")
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            var dataJson = JSON.parse(xhr.responseText)
            visitorsTable(dataJson);

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
                        labels: labels,
                        datasets: [{
                            label: 'DailyVisitors',
                            backgroundColor: 'rgb(14, 214, 24)',
                            borderColor: 'rgb(14, 214, 24)',
                            data: data,
                        }]
                    }
                });
        }
    };
    xhr.send(PassDate);

}
/*Generate Table */
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



}
/*Download Visitors Csv File calling Ajax file again*/
function VisitorsDownload() {
    FromDate = document.getElementById("fromDate").value;
    ToDate = document.getElementById("toDate").value;
    PassDate = "fromDate" + FromDate + "&toDate" + ToDate;
    var url = '/dailyUsers'

    const xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-url-urlencoded")
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            var dataJson = JSON.parse(xhr.responseText)
            const csvData =generateCsvVisitors(dataJson);
            download(csvData)

        }

    };
    xhr.send(PassDate);
}






const generateCsvVisitors= function (jsonCsv){
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
 function download(downloadData) {
    const blob = new Blob([downloadData],{type:'text/csv'});
    const url= URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.setAttribute('hidden','');
    a.setAttribute('href',url);
    a.setAttribute('download','visitors.csv');
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);

}
/*Dwelve time Visitors*/
function makeCharts2() {
    // FromDate = document.getElementById("fromDate").value;
    // ToDate = document.getElementById("toDate").value;
    // PassDate = "fromDate" + FromDate + "&toDate" + ToDate;
    var xhr = new XMLHttpRequest();
    xhr.open('POST','/avgDTime', true);
    xhr.onreadystatechange = function() {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200){
            var dataJson = JSON.parse(xhr.responseText)
            DwVisitorsTable(dataJson);

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
                        borderColor: 'rgb(133, 220, 32)',
                        fill:false,


                        data:data,
                    }]
                }
            });
        }
    };
    xhr.send(PassDate);

}
/*passing Dw data to table*/
function DwVisitorsTable(json) {
    console.log(json);

    //clear out existing table data
    while (tableDwVisitorsBody.firstChild) {
        tableDwVisitorsBody.removeChild(tableDwVisitorsBody.firstChild);
    }
//populate table
    // let matrix=[];
    json.forEach((test) => {
        tableDwVisitors.append("<tr>" +
            "<td class='col-xs-3'>" + "<h6>" + test.cpatured_date + "</h6></td> " +
            "<td class='col-xs-2'> <h6>" + test.captured_time_dw + "</h6>" + "</td>" +

            "</tr>");  });


}
/*Download Dw data to Csv file*/
function DwVisitorsDownload() {
    // FromDate = document.getElementById("fromDate").value;
    // ToDate = document.getElementById("toDate").value;
    // PassDate = "fromDate" + FromDate + "&toDate" + ToDate;
    var url = '/dailyUsers'

    const xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-url-urlencoded")
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            var dataJson = JSON.parse(xhr.responseText)
            const csvData =DwGenerateCsv(dataJson);
            downloadDwVisitors(csvData)

        }

    };
    xhr.send(PassDate);
}

function DwGenerateCsv (jsonCsv) {
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
 function downloadDwVisitors (downloadData) {
    const blob = new Blob([downloadData],{type:'text/csv'});
    const url= window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.setAttribute('hidden','');
    a.setAttribute('href',url);
    a.setAttribute('download','dwVisitors.csv');
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);

}

/*rpVistos Chart */

function makeCharts3() {
    // FromDate = document.getElementById("fromDate").value;
    // ToDate = document.getElementById("toDate").value;
    // PassDate = "fromDate" + FromDate + "&toDate" + ToDate;
    var xhr = new XMLHttpRequest();
    xhr.open('POST','/rpUsers', true);
    xhr.onreadystatechange = function() {
        if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200){
            var dataJson = JSON.parse(xhr.responseText)
            rpVisitorsTable(dataJson)

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
                        label: 'RpVisitors',
                        backgroundColor: 'rgb( 26, 249, 112)',
                        borderColor: 'rgb( 29, 49, 7)',
                        data:data,
                    }]            }
            });
        }
    };
    xhr.send(PassDate);

}

/*rp user table*/
function rpVisitorsTable(json) {
    console.log(json);

    //clear out existing table data
    while (tableRpVisitorsBody.firstChild) {
        tableRpVisitorsBody.removeChild(tableRpVisitorsBody.firstChild);
    }
//populate table
    // let matrix=[];
    json.forEach((test) => {
        tableRpVisitors.append("<tr>" +
            "<td class='col-xs-3'>" + "<h6>" + test.p_date + "</h6></td> " +
            "<td class='col-xs-2'> <h6>" + test.rp_count + "</h6>" + "</td>" +

            "</tr>");  });
    //RpDownloadToCsv(json)


}
/*Download Rp Visitors to Csv file*/
function RpVisitorsDownload() {
    FromDate = document.getElementById("fromDate").value;
    ToDate = document.getElementById("toDate").value;
    PassDate = "fromDate" + FromDate + "&toDate" + ToDate;
    var url = '/dailyUsers'

    const xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("Content-type", "application/x-www-form-url-urlencoded")
    xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200) {
            var dataJson = JSON.parse(xhr.responseText)
            const csvData =RpGenerateCsv(dataJson);
            RpDownload(csvData)

        }

    };
    xhr.send(PassDate);
}

const RpGenerateCsv=function (jsonCsv) {
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
const RpDownload = function (downloadData) {
    const blob = new Blob([downloadData],{type:'text/csv'});
    const url= window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.setAttribute('hidden','');
    a.setAttribute('href',url);
    a.setAttribute('download','rpVisitors.csv');
    document.body.appendChild(a);
    a.click();
    document.body.removeChild(a);

}


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


//login page check
//var formUse =document.querySelector('#form-user-login')
//var  loginsubmit =document.querySelector('#loginsubmit')
var loginError = document.querySelector('#login-error')

function checkLoginStatus(){
    var xhr = new XMLHttpRequest();
    xhr.open(('post','/login',true))
    xhr.addEventListener('readystatechange',function () {
        if(xhr.readyState == XMLHttpRequest.DONE && xhr.status == 200){
            var item = xhr.responseText;
            console.log(item);
            if (item == 'false'){
                loginError.textContent='invalid usename or password'
            }else {
                loginError.textContent = '';

            }
            
        }


    });
    xhr.send()

}

