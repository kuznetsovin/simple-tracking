create table if not exists geo_object
(
    id   SERIAL PRIMARY KEY,
    name varchar(10)                  not null,
    geom geometry(MultiPolygon, 3857) not null
);

comment on column geo_object.name is 'Название объекта';
comment on column geo_object.geom is 'Геометрия';

CREATE UNIQUE INDEX geo_object__index ON geo_object USING btree (name);