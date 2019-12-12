create schema if not exists last_point;
create table if not exists last_point.clients
(
    id            int CONSTRAINT last_point_client_uniq UNIQUE,
    packet_id     bigint,
    navigate_date timestamp,
    speed         double precision,
    nsat          smallint,
    course        smallint,
    pdop          int,
    geom          geometry(Point, 4326),
    update_date   timestamp without time zone
);
comment on column last_point.clients.id is 'код бнсо';
comment on column last_point.clients.packet_id is 'идентификатор пакета нот устройства';
comment on column last_point.clients.navigate_date is 'дата точки';
comment on column last_point.clients.navigate_date is 'дата приема';
comment on column last_point.clients.speed is 'скорость';
comment on column last_point.clients.nsat is 'кол-во спутников';
comment on column last_point.clients.course is 'направление';
comment on column last_point.clients.pdop is 'коэффициент показателя точности';
comment on column last_point.clients.geom is 'геометрия';

