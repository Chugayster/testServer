--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

-- Started on 2022-05-26 19:24:00

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 3 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3313 (class 0 OID 0)
-- Dependencies: 3
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 210 (class 1259 OID 16415)
-- Name: events; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.events (
    id integer NOT NULL,
    name text,
    location text
);


ALTER TABLE public.events OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16414)
-- Name: new_table_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.new_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.new_table_id_seq OWNER TO postgres;

--
-- TOC entry 3314 (class 0 OID 0)
-- Dependencies: 209
-- Name: new_table_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.new_table_id_seq OWNED BY public.events.id;


--
-- TOC entry 3164 (class 2604 OID 16418)
-- Name: events id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events ALTER COLUMN id SET DEFAULT nextval('public.new_table_id_seq'::regclass);


--
-- TOC entry 3307 (class 0 OID 16415)
-- Dependencies: 210
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.events (id, name, location) FROM stdin;
5	Pool party	Pine forest
6	Loft party	Mariupolska Str
7	New Year party	Central square
3	Forest party	Pine forest
2	Pool party	Public pool
1	Birthday party	98 str
4	Sea party	Azov Sea
18	Food festival	Central square
19	Independence day	City hall
24	Victory Day	Victory square
23	City day	City garden
\.


--
-- TOC entry 3315 (class 0 OID 0)
-- Dependencies: 209
-- Name: new_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.new_table_id_seq', 26, true);


--
-- TOC entry 3166 (class 2606 OID 16422)
-- Name: events new_table_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT new_table_pkey PRIMARY KEY (id);


-- Completed on 2022-05-26 19:24:00

--
-- PostgreSQL database dump complete
--

