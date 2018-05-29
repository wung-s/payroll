## Project Description

Employer pays its employees by the hour (there are no
salaried employees.) Employees belong to one of two _job groups_ which
determine their wages; job group A is paid $20/hr, and job group B is paid
$30/hr. Each employee is identified by a string called an "employee id" that is
globally unique in their system.

Hours are tracked per employee, per day in comma-separated value files (CSV).
Each individual CSV file is known as a "time report", and will contain:

1.  A header, denoting the columns in the sheet (`date`, `hours worked`,
    `employee id`, `job group`)
1.  0 or more data rows
1.  A footer row where the first cell contains the string `report id`, and the
    second cell contains a unique identifier for this report.

File content guarantees:

1.  Columns will always be in that order.
1.  There will always be data in each column.
1.  There will always be a well-formed header line.
1.  There will always be a well-formed footer line.

### Important:

Example input files named `sample.csv` and `sample2.csv` are included in this repo.

### What the web-based application does:

We've agreed to build the following web-based prototype for our partner.

1.  Accepts a comma separated file with the schema described in the previous section.
1.  Parses the given file, and store the timekeeping information in
    a relational database for archival reasons.
1.  After upload, the application should displays a _payroll report_. This
    report is accessible to the user without them having to upload a
    file first.
1.  If an attempt is made to upload two files with the same report id, the
    second upload fails

The payroll report is structured as follows:

1.  There are 3 columns in the report: `Employee Id`, `Pay Period`,
    `Amount Paid`
1.  A `Pay Period` is a date interval that is roughly biweekly. Each month has
    two pay periods; the _first half_ is from the 1st to the 15th inclusive, and
    the _second half_ is from the 16th to the end of the month, inclusive.
1.  Each employee will have a single row in the report for each pay period
    that they have recorded hours worked. The `Amount Paid` is be reported
    as the sum of the hours worked in that pay period multiplied by the hourly
    rate for their job group.
1.  If an employee was not paid in a specific pay period, there will not be a
    row for that employee + pay period combination in the report.
1.  The report should be sorted by employee id and then pay period start
1.  The report is based on all _of the data_ across _all of the uploaded
    time reports_, for all time.

As an example, a sample file with the following data:

<table>
<tr>
  <th>
    date
  </th>
  <th>
    hours worked
  </th>
  <th>
    employee id
  </th>
  <th>
    job group
  </th>
</tr>
<tr>
  <td>
    4/11/2016
  </td>
  <td>
    10
  </td>
  <td>
    1
  </td>
  <td>
    A
  </td>
</tr>
<tr>
  <td>
    14/11/2016
  </td>
  <td>
    5
  </td>
  <td>
    1
  </td>
  <td>
    A
  </td>
</tr>
<tr>
  <td>
    20/11/2016
  </td>
  <td>
    3
  </td>
  <td>
    2
  </td>
  <td>
    B
  </td>
</tr>
</table>

should produce the following payroll report:

<table>
<tr>
  <th>
    Employee ID
  </th>
  <th>
    Pay Period
  </th>
  <th>
    Amount Paid
  </th>
</tr>
<tr>
  <td>
    1
  </td>
  <td>
    1/11/2016 - 15/11/2016
  </td>
  <td>
    $300.00
  </td>
</tr>
  <td>
    2
  </td>
  <td>
    16/11/2016 - 30/11/2016
  </td>
  <td>
    $90.00
  </td>
</tr>
</table>

### Set Up

    $ mkdir -p $GOPATH/src/github.com/wung-s
    $ cd $GOPATH/src/github.com/wung-s
    $ git clone git@github.com:wung-s/payroll.git && cd payroll
    $ buffalo db create -a
    $ buffalo db migrate
    $ buffalo task db:seed

## Starting the Application

    $ PORT=4000 buffalo dev

## Environment Variable

    TIMEZONE=EDT

## Assumptions

* Only one partner
* No authentication
* Date in the `csv` are always in the format `DD/MM/YYYY`

[Powered by Buffalo](http://gobuffalo.io)
