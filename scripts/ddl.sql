DROP TABLE IF EXISTS public.flights;
DROP TABLE IF EXISTS public.flights_segments;
DROP TABLE IF EXISTS public.iata;

CREATE TABLE IF NOT EXISTS public.iata
(
    id integer GENERATED ALWAYS AS IDENTITY,
    name character varying(100),
    code character varying(5),
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS public.flights_segments
(
    id integer GENERATED ALWAYS AS IDENTITY,
    name character varying(100),
    origin character varying(100),
    destination character varying(100),
    miles integer,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS public.flights
(
    id integer GENERATED ALWAYS AS IDENTITY,
    segment_id integer NOT NULL,
    scheduled_departure_time timestamptz,
    scheduled_arrival_time timestamptz,
    first_class_base_cost decimal,
    economy_class_base_cost decimal,
    num_first_class_seats int,
    num_economy_class_seats int,
    airplane_type_id character varying(100),
    PRIMARY KEY(id),
    UNIQUE(id, segment_id),
    CONSTRAINT fk_flight_segment
      FOREIGN KEY(segment_id) 
	  REFERENCES flights_segments(id)
);

