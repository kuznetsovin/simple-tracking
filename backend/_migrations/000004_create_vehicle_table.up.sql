create table if not exists vehicle
(
    id         SERIAL PRIMARY KEY,
    gos_number varchar(10) not null,
    gps_code   int         not null
);

comment on column vehicle.gos_number is 'Гос. номер';
comment on column vehicle.gps_code is 'Код навигационного устроства';

CREATE UNIQUE INDEX vehicle__index ON vehicle USING btree (gos_number, gps_code);