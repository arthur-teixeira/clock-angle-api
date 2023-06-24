--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3 (Debian 15.3-1.pgdg120+1)
-- Dumped by pg_dump version 15.3 (Debian 15.3-1.pgdg120+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: angles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.angles (
    id integer NOT NULL,
    angle integer NOT NULL,
    hour integer NOT NULL,
    minute integer NOT NULL,
    date date NOT NULL
);


ALTER TABLE public.angles OWNER TO postgres;

--
-- Name: angles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.angles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.angles_id_seq OWNER TO postgres;

--
-- Name: angles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.angles_id_seq OWNED BY public.angles.id;


--
-- Name: angles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.angles ALTER COLUMN id SET DEFAULT nextval('public.angles_id_seq'::regclass);


--
-- Data for Name: angles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.angles (id, angle, hour, minute, date) FROM stdin;
1	140	12	40	2023-06-24
2	113	12	45	2023-06-24
3	0	12	0	2023-06-24
4	30	13	0	2023-06-24
5	135	13	30	2023-06-24
\.


--
-- Name: angles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.angles_id_seq', 5, true);


--
-- Name: angles angles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.angles
    ADD CONSTRAINT angles_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

