$(function() {
    $("#customer_reportrange").daterangepicker({

        autoUpdateInput: false,
        locale: {
            cancelLabel: 'Clear'
        },
        ranges: {
            // 'Today': [moment(), moment()],
            'Yesterday': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
            'Last 7 Days': [moment().subtract(6, 'days'), moment()],
            'Last 30 Days': [moment().subtract(29, 'days'), moment()],
            'This Month': [moment().startOf('month'), moment().endOf('month')],
            'Last Month': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')]
        }
    });

    $("#customer_reportrange").on('apply.daterangepicker', function(ev, picker) {
        $(this).val(picker.startDate.format('MM/DD/YYYY') + ' - ' + picker.endDate.format('MM/DD/YYYY'));



        var url = window.location.search.substring(1);
        var parsed_qs = parse_query_string(url);
        //	console.log(url);

        //alert(picker.startDate.format('MM/DD/YYYY'));
        globalStartDate = picker.startDate.format('YYYY-MM-DD');
        globalEndDate = picker.endDate.format('YYYY-MM-DD');


        var url = window.location.search.substring(1);
        var parsed_qs = parse_query_string(url);
        //	console.log(url);

        //alert(picker.startDate.format('MM/DD/YYYY'));
        globalStartDate = picker.startDate.format('YYYY-MM-DD');
        globalEndDate = picker.endDate.format('YYYY-MM-DD');

        LoadCustomerData(globalStartDate,globalEndDate,InitialCustomerList);
        //LoadCustomerData(globalStartDate,globalEndDate,parsed_qs.customerID);


    });

        $("#customer_reportrange").on('cancel.daterangepicker', function(ev, picker) {
            $(this).val('');
        });




    }

);
$('#showDateRange').text(StartDate + " To " + EndDate);