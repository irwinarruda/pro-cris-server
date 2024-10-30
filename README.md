# Pro Cris Server

## Steps

- 1 Install [air](https://github.com/cosmtrek/air)
- 2 Install [templ](https://templ.guide/quick-start/installation)
- 3 Use `make dev` to run the server in development mode.
- 4 Use `make test` to start runing unit tests.

## Endpoints

- GET /api/v1/checkhealth
- POST /api/v1/auth/login
- Middleware /api/v1/auth/ensure-authenticated
- POST /api/v1/student
- PUT /api/v1/student/:id
- PUT /api/v1/student/:id/picture
- GET /api/v1/student
- GET /api/v1/student/:id
- DELETE /api/v1/student/:id
- POST /api/v1/appointment
- PUT /api/v1/appointment/:id
- GET /api/v1/appointment
- GET /api/v1/appointment/:id
- DELETE /api/v1/appointment/:id

## Similar apps

- https://iprofe.com.br/
- https://www.iorclass.com.br/?af=deprof

## Student

- Get All Students.

  - Check fields to be sent there

- Get Student by ID.

  - Check fields to be sent there

- Create Student with

  - Base information
  - PaymentMethod (**Upfront** or **Later**)
  - PaymentStyle (**Fixed** or **Variable**)
  - PaymentStyleValue
  - SettlementMethod (**NumberAppointments** or **AmountOfTime** or **None**)
  - SettlementMethodStartDate
  - SettlementMethodValue

- Update Student and it's default information.
- Delete Student.

## Appointments

- Get appointments and future appointments by date.
- Create daily appointments by all students routine.
- Easily cancel an appointment right before or right after routine created.
- Create manual appointment for a student. Flag if it's extra class or not

- Get all student appointments.
  - Filter by paid and not paid.
  - Filter by date.
  - Filter by settled and not settled.
  - Filter by canceled and not canceled.
- Show all student not settled appointments.
  - have recipt sent to student WhatsApp.
  - Show sent recipts.
  - possibility to edit price when sending
- Notification for when you should send a recipt to a student.

## Settlement

- Get student settlement type
- Get amount already paid by month (filter)

- Start Settlement based on params;
- Stop Settlement based on params;
- Mark Settlement as paid;
- Generate prediction based on **Routine** and **SettlementMethod**;
- See current Settlement statitics based on SettlementMethod;
- See current Settlement prediction;
- Store amount paid and when it was paid;

- SettlementMethod == AmountOfTime

  - Start receipt count based on SettlementMethodStartDate;
  - Stop receipt count if SettlementMethodStartDate + SettlementMethodValue was reached;
  - Start receipt count if no SettlementMethodStartDate;
  - Stop receipt count if no SettlementMethodStartDate;
  - Notificação indicating if the SettlementMethodStartDate is close to be achieved;

- SettlementMethod == NumberAppointments

  - Start receipt count based on SettlementMethodStartDate;
  - Stop receipt count if SettlementMethodValue was reached;
  - Start receipt count if no SettlementMethodStartDate;
  - Stop receipt count if no SettlementMethodStartDate;
  - Notificação indicating if the SettlementMethodValue is close to be achieved;

- SettlementMethod == NumberAppointments && SettlementMethod == AmountOfTime
