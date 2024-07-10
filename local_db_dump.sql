--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3 (Ubuntu 16.3-1.pgdg22.04+1)
-- Dumped by pg_dump version 16.3 (Ubuntu 16.3-1.pgdg22.04+1)

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
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: carts; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.carts (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    transaction_id bigint
);


ALTER TABLE public.carts OWNER TO postgres;

--
-- Name: carts_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.carts ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME public.carts_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: files; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.files (
    id bigint NOT NULL,
    inserted_at timestamp(0) without time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    updated_at timestamp(0) without time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    user_id bigint,
    local_path character varying(255),
    file_name character varying(255),
    technology character varying(255),
    color character varying(255)
);


ALTER TABLE public.files OWNER TO postgres;

--
-- Name: files_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.files ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME public.files_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: goose_db_version; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);


ALTER TABLE public.goose_db_version OWNER TO postgres;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.goose_db_version_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.goose_db_version_id_seq OWNER TO postgres;

--
-- Name: goose_db_version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.goose_db_version_id_seq OWNED BY public.goose_db_version.id;


--
-- Name: orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.orders (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    file_id bigint NOT NULL,
    inserted_at timestamp(0) without time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    updated_at timestamp(0) without time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    build_time bigint NOT NULL,
    quantity character varying(255) NOT NULL,
    price double precision NOT NULL,
    payment_complete boolean NOT NULL,
    status character varying(255) NOT NULL,
    print_status character varying(50),
    printer_id bigint
);


ALTER TABLE public.orders OWNER TO postgres;

--
-- Name: orders_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.orders ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME public.orders_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: transaction_orders; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transaction_orders (
    transaction_id bigint NOT NULL,
    order_id bigint NOT NULL
);


ALTER TABLE public.transaction_orders OWNER TO postgres;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.transactions (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    payment_status character varying(255) NOT NULL,
    checkout_request_id character varying(255) NOT NULL,
    inserted_at timestamp(0) without time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    updated_at timestamp(0) without time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    phone character varying(20)
);


ALTER TABLE public.transactions OWNER TO postgres;

--
-- Name: transactions_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.transactions ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME public.transactions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    email public.citext NOT NULL,
    password_hash character varying(255),
    inserted_at timestamp(0) without time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    updated_at timestamp(0) without time zone DEFAULT (now() AT TIME ZONE 'utc'::text) NOT NULL,
    has_printer boolean DEFAULT false
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: goose_db_version id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goose_db_version ALTER COLUMN id SET DEFAULT nextval('public.goose_db_version_id_seq'::regclass);


--
-- Data for Name: carts; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.carts (id, user_id, transaction_id) FROM stdin;
4	98	\N
\.


--
-- Data for Name: files; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.files (id, inserted_at, updated_at, user_id, local_path, file_name, technology, color) FROM stdin;
58	2024-07-04 09:01:16	2024-07-04 09:01:16	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
59	2024-07-04 09:02:52	2024-07-04 09:02:52	98	left_rod_mount.stl	left_rod_mount.stl	FDM (Plastic)	Any
60	2024-07-04 09:03:13	2024-07-04 09:03:13	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
61	2024-07-04 09:03:35	2024-07-04 09:03:35	98	left_rod_mount.stl	left_rod_mount.stl	FDM (Plastic)	Any
62	2024-07-04 09:05:06	2024-07-04 09:05:06	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
63	2024-07-04 09:07:17	2024-07-04 09:07:17	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
64	2024-07-05 11:14:54	2024-07-05 11:14:54	98	left_rod_mount.stl	left_rod_mount.stl	FDM (Plastic)	Any
65	2024-07-05 11:22:29	2024-07-05 11:22:29	98	left_rod_mount.stl	left_rod_mount.stl	FDM (Plastic)	Any
66	2024-07-05 11:28:12	2024-07-05 11:28:12	98	left_rod_mount.stl	left_rod_mount.stl	FDM (Plastic)	Any
67	2024-07-05 11:33:34	2024-07-05 11:33:34	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
68	2024-07-05 11:47:02	2024-07-05 11:47:02	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
69	2024-07-05 11:51:58	2024-07-05 11:51:58	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
70	2024-07-05 11:56:59	2024-07-05 11:56:59	98	left_rod_mount.stl	left_rod_mount.stl	FDM (Plastic)	Any
71	2024-07-07 21:38:31	2024-07-07 21:38:31	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
72	2024-07-08 10:34:25	2024-07-08 10:34:25	98	left_rod_mount.stl	left_rod_mount.stl	FDM (Plastic)	Any
73	2024-07-08 10:51:52	2024-07-08 10:51:52	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
74	2024-07-09 08:58:05	2024-07-09 08:58:05	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
75	2024-07-09 08:58:39	2024-07-09 08:58:39	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
76	2024-07-09 09:00:00	2024-07-09 09:00:00	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
77	2024-07-09 09:19:53	2024-07-09 09:19:53	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
78	2024-07-09 09:21:14	2024-07-09 09:21:14	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
79	2024-07-09 09:22:31	2024-07-09 09:22:31	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
80	2024-07-09 09:24:54	2024-07-09 09:24:54	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
81	2024-07-09 09:35:19	2024-07-09 09:35:19	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
82	2024-07-09 09:36:46	2024-07-09 09:36:46	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
83	2024-07-09 09:38:21	2024-07-09 09:38:21	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
84	2024-07-09 09:40:02	2024-07-09 09:40:02	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
85	2024-07-09 14:51:06	2024-07-09 14:51:06	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
86	2024-07-09 18:28:00	2024-07-09 18:28:00	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
87	2024-07-09 18:29:54	2024-07-09 18:29:54	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
88	2024-07-09 18:30:59	2024-07-09 18:30:59	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
89	2024-07-09 18:32:37	2024-07-09 18:32:37	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
90	2024-07-09 19:11:36	2024-07-09 19:11:36	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
91	2024-07-09 19:14:31	2024-07-09 19:14:31	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
92	2024-07-09 19:16:09	2024-07-09 19:16:09	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
93	2024-07-09 19:37:26	2024-07-09 19:37:26	98	ServoArm.STL	ServoArm.STL	FDM (Plastic)	Any
\.


--
-- Data for Name: goose_db_version; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.goose_db_version (id, version_id, is_applied, tstamp) FROM stdin;
1	0	t	2024-06-27 11:10:38.89145
2	20240627075415	t	2024-06-27 11:12:09.697316
3	20240630073818	t	2024-06-30 11:03:43.557615
4	20240630085100	t	2024-06-30 12:02:59.460434
5	20240630093424	t	2024-06-30 12:36:15.637163
6	20240702075012	t	2024-07-02 12:42:06.734347
7	20240702094648	t	2024-07-02 12:50:02.08995
8	20240702103013	t	2024-07-02 13:32:03.234856
9	20240703063440	t	2024-07-03 09:38:42.454966
10	20240703082152	t	2024-07-03 11:24:32.226957
11	20240704120540	t	2024-07-04 15:10:25.864631
12	20240705082003	t	2024-07-05 11:21:55.644525
13	20240705082326	t	2024-07-05 11:27:43.381663
14	20240705083029	t	2024-07-05 11:32:54.70819
15	20240709133614	t	2024-07-09 17:11:34.399714
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (id, user_id, file_id, inserted_at, updated_at, build_time, quantity, price, payment_complete, status, print_status, printer_id) FROM stdin;
38	98	73	2024-07-08 10:51:52	2024-07-08 10:51:55	72	1	46.736343011271096	f	Reviewing	Available	\N
39	98	74	2024-07-09 08:58:05	2024-07-09 08:58:08	72	1	46.736343011271096	f	Reviewing	Available	\N
40	98	75	2024-07-09 08:58:39	2024-07-09 08:58:42	72	1	46.736343011271096	f	Reviewing	Available	\N
41	98	76	2024-07-09 09:00:00	2024-07-09 09:00:05	72	1	46.736343011271096	f	Reviewing	Available	\N
42	98	77	2024-07-09 09:19:53	2024-07-09 09:19:56	72	1	46.736343011271096	f	Reviewing	Available	\N
43	98	78	2024-07-09 09:21:14	2024-07-09 09:21:17	72	1	46.736343011271096	f	Reviewing	Available	\N
44	98	79	2024-07-09 09:22:31	2024-07-09 09:22:37	72	1	46.736343011271096	f	Reviewing	Available	\N
45	98	80	2024-07-09 09:24:54	2024-07-09 09:24:57	72	1	46.736343011271096	f	Reviewing	Available	\N
46	98	81	2024-07-09 09:35:19	2024-07-09 09:35:22	72	1	46.736343011271096	f	Reviewing	Available	\N
47	98	82	2024-07-09 09:36:46	2024-07-09 09:36:49	72	1	46.736343011271096	f	Reviewing	Available	\N
34	98	69	2024-07-05 11:51:58	2024-07-08 09:55:49	72	1	46.736343011271096	f	Completed	Completed	98
48	98	83	2024-07-09 09:38:21	2024-07-09 09:38:24	72	1	46.736343011271096	f	Reviewing	Available	\N
49	98	84	2024-07-09 09:40:02	2024-07-09 09:40:12	72	1	46.736343011271096	f	Reviewing	Available	\N
50	98	85	2024-07-09 14:51:06	2024-07-09 14:51:09	72	1	46.736343011271096	f	Reviewing	Available	\N
35	98	70	2024-07-05 11:56:59	2024-07-08 10:03:12	72	1	1223.457307794972	f	Completed	Completed	98
8	98	63	2024-07-04 09:07:17	2024-07-08 10:03:13	72	1	46.736343011271096	f	Completed	Completed	98
51	98	86	2024-07-09 18:28:00	2024-07-09 18:28:05	72	1	46.736343011271096	f	Reviewing	Available	\N
27	98	58	2024-07-04 09:01:16	2024-07-08 10:20:36	72	1	46.736343011271096	f	Completed	Completed	98
28	98	63	2024-07-04 09:07:17	2024-07-08 10:22:17	72	1	46.736343011271096	f	Completed	Completed	98
52	98	87	2024-07-09 18:29:54	2024-07-09 18:31:02	72	1	46.736343011271096	f	Reviewing	Available	\N
53	98	88	2024-07-09 18:30:59	2024-07-09 18:31:02	72	1	46.736343011271096	f	Reviewing	Available	\N
54	98	89	2024-07-09 18:32:37	2024-07-09 18:32:40	72	1	46.736343011271096	f	Reviewing	Available	\N
55	98	90	2024-07-09 19:11:36	2024-07-09 19:11:36	72	1	46.736343011271096	f	Reviewing	Available	\N
56	98	91	2024-07-09 19:14:31	2024-07-09 19:14:31	72	1	46.736343011271096	f	Reviewing	Available	\N
57	98	92	2024-07-09 19:16:09	2024-07-09 19:16:09	72	1	46.736343011271096	f	Reviewing	Available	\N
58	98	93	2024-07-09 19:37:26	2024-07-09 19:37:26	72	1	46.736343011271096	f	Reviewing	Available	\N
36	98	71	2024-07-07 21:38:31	2024-07-08 10:29:10	72	1	46.736343011271096	f	Completed	Completed	98
33	98	68	2024-07-05 11:47:02	2024-07-08 10:29:15	72	1	46.736343011271096	f	Completed	Completed	98
37	98	72	2024-07-08 10:34:25	2024-07-08 10:37:57	72	1	1223.457307794972	f	Completed	Completed	98
\.


--
-- Data for Name: transaction_orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transaction_orders (transaction_id, order_id) FROM stdin;
2	52
2	53
3	54
4	55
4	56
4	57
4	58
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transactions (id, user_id, payment_status, checkout_request_id, inserted_at, updated_at, phone) FROM stdin;
2	98	AwaitingPayment	ws_CO_09072024183429674713958070	2024-07-09 15:31:04	2024-07-09 15:31:04	\N
3	98	AwaitingPayment	ws_CO_09072024183608599713958070	2024-07-09 15:32:42	2024-07-09 15:32:42	\N
4	98	ProcessingPayment	ws_CO_09072024193735890713958070	2024-07-09 16:11:36	2024-07-09 16:11:36	0713958070
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.users (id, email, password_hash, inserted_at, updated_at, has_printer) FROM stdin;
99	mwaurad@mail.com	$2a$10$GebcvV9OOSvnRPe2kQ75dO5XYfYT8wFR9gyyIJ4L5nHRdpJh5yn/6	2024-07-03 12:04:25	2024-07-03 12:04:25	f
98	mwaurathealex@gmail.com	$2a$10$cjejOu6tHT.UHid2kIjSsOv1GHPsrGNcjIJc4/5KyPsgFDGyldsGC	2024-07-03 12:03:35	2024-07-03 12:03:35	t
\.


--
-- Name: carts_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.carts_id_seq', 4, true);


--
-- Name: files_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.files_id_seq', 93, true);


--
-- Name: goose_db_version_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.goose_db_version_id_seq', 15, true);


--
-- Name: orders_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_id_seq', 58, true);


--
-- Name: transactions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_id_seq', 4, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.users_id_seq', 99, true);


--
-- Name: carts carts_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT carts_pkey PRIMARY KEY (id);


--
-- Name: files files_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_pkey PRIMARY KEY (id);


--
-- Name: goose_db_version goose_db_version_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.goose_db_version
    ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: transaction_orders transaction_orders_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_orders
    ADD CONSTRAINT transaction_orders_pkey PRIMARY KEY (transaction_id, order_id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: files files_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.files
    ADD CONSTRAINT files_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: orders fk_file; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_file FOREIGN KEY (file_id) REFERENCES public.files(id) ON DELETE CASCADE;


--
-- Name: orders fk_printer; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_printer FOREIGN KEY (printer_id) REFERENCES public.users(id) ON DELETE SET NULL;


--
-- Name: carts fk_transactions; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT fk_transactions FOREIGN KEY (transaction_id) REFERENCES public.transactions(id) ON DELETE CASCADE;


--
-- Name: orders fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: carts fk_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.carts
    ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: transaction_orders transaction_orders_order_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_orders
    ADD CONSTRAINT transaction_orders_order_id_fkey FOREIGN KEY (order_id) REFERENCES public.orders(id) ON DELETE CASCADE;


--
-- Name: transaction_orders transaction_orders_transaction_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transaction_orders
    ADD CONSTRAINT transaction_orders_transaction_id_fkey FOREIGN KEY (transaction_id) REFERENCES public.transactions(id) ON DELETE CASCADE;


--
-- Name: transactions transactions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

