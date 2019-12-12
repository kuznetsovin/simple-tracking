create schema if not exists vts;
create table if not exists vts.track
(
  client         int,
  packet_id      bigint,
  navigate_date  timestamp,
  received_date  timestamp,
  longitude      double precision,
  latitude       double precision,
  speed          double precision,
  nsat           smallint,
  course         smallint,
  pdop           int,
  geom           geometry(Point, 4326),
  kbm_sensors    jsonb,
  liquid_sensors jsonb
)
partition by range (navigate_date);
-- CREATE INDEX client_nav_date__index ON vts.track USING btree (client, navigate_date);

create table vts.track_201911 partition of vts.track
    for values from ('2019-11-01') to ('2019-11-30 23:59:59.999');

create table vts.track_201912 partition of vts.track
    for values from ('2019-12-01') to ('2019-12-31 23:59:59.999');

comment on column vts.track.client is 'код бнсо';
comment on column vts.track.packet_id is 'идентификатор пакета нот устройства';
comment on column vts.track.navigate_date is 'дата точки';
comment on column vts.track.navigate_date is 'дата приема';
comment on column vts.track.longitude is 'долгота';
comment on column vts.track.latitude is 'широта';
comment on column vts.track.speed is 'скорость';
comment on column vts.track.nsat is 'кол-во спутников';
comment on column vts.track.course is 'направление';
comment on column vts.track.pdop is 'коэффициент показателя точности';
comment on column vts.track.liquid_sensors is 'json по дутам';
comment on column vts.track.kbm_sensors is 'json по кбмам';
comment on column vts.track.geom is 'геометрия';
