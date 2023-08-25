--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4 (Ubuntu 15.4-0ubuntu0.23.04.1)
-- Dumped by pg_dump version 15.4 (Ubuntu 15.4-0ubuntu0.23.04.1)

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

DROP DATABASE IF EXISTS calendar;
--
-- Name: calendar; Type: DATABASE; Schema: -; Owner: calendar
--

CREATE DATABASE calendar WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'C';

CREATE USER calendar PASSWORD 'pass';

ALTER DATABASE calendar OWNER TO calendar;

\connect calendar

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
-- Name: events; Type: TABLE; Schema: public; Owner: calendar
--

CREATE TABLE public.events (
    event_id character(36) NOT NULL,
    title character varying(255) NOT NULL,
    description text,
    date_time_start timestamp with time zone NOT NULL,
    date_time_end timestamp with time zone NOT NULL,
    date_time_notice timestamp with time zone NOT NULL,
    user_id character(36) NOT NULL
);


ALTER TABLE public.events OWNER TO calendar;

--
-- Name: notes_check; Type: TABLE; Schema: public; Owner: calendar
--

CREATE TABLE public.notes_check (
    last_check_date_time timestamp with time zone NOT NULL
);


ALTER TABLE public.notes_check OWNER TO calendar;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: calendar
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO calendar;

--
-- Data for Name: events; Type: TABLE DATA; Schema: public; Owner: calendar
--

INSERT INTO public.events VALUES ('de0fe01f-96a5-b53c-2cbb-12f9b00099d0', 'Title Notice 26', 'Description Notice 26', '2023-08-20 21:07:18+03', '2023-08-20 21:27:18+03', '2023-08-20 20:57:23+03', 'ea51c062-ede7-ed54-fd25-0525ddc8d25a');
INSERT INTO public.events VALUES ('2a73425d-82d6-a340-f8b0-a0574b2998da', 'Title Notice 28', 'Description Notice 28', '2023-08-20 21:07:19+03', '2023-08-20 21:27:19+03', '2023-08-20 20:57:24+03', 'e7adb469-9447-e9a7-81d5-67c42a8f247a');
INSERT INTO public.events VALUES ('948010b7-ccd0-995d-c58a-b29975ea853c', 'Title Notice 2', 'Description Notice 2', '2023-08-20 21:07:06+03', '2023-08-20 21:27:06+03', '2023-08-20 20:57:11+03', 'c2fefdb5-95fe-a1a4-3f94-43f99326bad7');
INSERT INTO public.events VALUES ('53ee3846-b9e1-07be-f473-a525d6e34dba', 'Title Notice 4', 'Description Notice 4', '2023-08-20 21:07:07+03', '2023-08-20 21:27:07+03', '2023-08-20 20:57:12+03', '6df46399-e1d2-4eeb-6b34-198880f594e6');
INSERT INTO public.events VALUES ('c3b12c7b-f6d2-b65d-5cf3-f74a96c5ff8d', 'Title Notice 6', 'Description Notice 6', '2023-08-20 21:07:08+03', '2023-08-20 21:27:08+03', '2023-08-20 20:57:13+03', '091e35b7-51e6-f7c5-1828-c9ee0a13f2b3');
INSERT INTO public.events VALUES ('d1bea351-5f64-6c98-1840-3d9d9c6c8ed1', 'Title Notice 8', 'Description Notice 8', '2023-08-20 21:07:09+03', '2023-08-20 21:27:09+03', '2023-08-20 20:57:14+03', '6b347fba-d26c-94cf-3033-70dd57c031bb');
INSERT INTO public.events VALUES ('84b4fec3-8007-6c89-75e7-5f142e8b4aac', 'Title Notice 10', 'Description Notice 10', '2023-08-20 21:07:10+03', '2023-08-20 21:27:10+03', '2023-08-20 20:57:15+03', 'c6ac5c06-a3b9-642a-fdfb-c34a0902f602');
INSERT INTO public.events VALUES ('537377ee-d00a-b219-078c-03915b33d197', 'Title Notice 12', 'Description Notice 12', '2023-08-20 21:07:11+03', '2023-08-20 21:27:11+03', '2023-08-20 20:57:16+03', '2325956a-5094-d532-8dd8-d94429d162b5');
INSERT INTO public.events VALUES ('501be5ad-e715-b1f9-3410-320c221e8b69', 'Title Notice 14', 'Description Notice 14', '2023-08-20 21:07:12+03', '2023-08-20 21:27:12+03', '2023-08-20 20:57:17+03', '3b993ef1-3301-7989-f58e-6bc43cf12cb9');
INSERT INTO public.events VALUES ('dfd2a488-77ac-bb4f-6100-f050eda07150', 'Title Notice 16', 'Description Notice 16', '2023-08-20 21:07:13+03', '2023-08-20 21:27:13+03', '2023-08-20 20:57:18+03', 'ed253608-4629-2e67-5b1b-cccdc0d0c03c');
INSERT INTO public.events VALUES ('b0e6a28e-f755-8944-b319-9df95a44a5b0', 'Title Notice 18', 'Description Notice 18', '2023-08-20 21:07:14+03', '2023-08-20 21:27:14+03', '2023-08-20 20:57:19+03', 'ea7374bb-47e7-346e-bbef-6e8e44f9ce9f');
INSERT INTO public.events VALUES ('3d13db6f-5efd-b3c3-f5eb-c40e6ad8cbdb', 'Title Notice 20', 'Description Notice 20', '2023-08-20 21:07:15+03', '2023-08-20 21:27:15+03', '2023-08-20 20:57:20+03', 'a123401a-689f-dd30-1007-168ee9a02eda');
INSERT INTO public.events VALUES ('67c873b5-2ca3-1263-60f9-f94d2272dcb6', 'Title Notice 22', 'Description Notice 22', '2023-08-20 21:07:16+03', '2023-08-20 21:27:16+03', '2023-08-20 20:57:21+03', 'f125d385-b583-10b2-8ab1-e3a42ce67127');
INSERT INTO public.events VALUES ('0a3bebc8-a418-66f8-232a-01791815225e', 'Title Notice 24', 'Description Notice 24', '2023-08-20 21:07:17+03', '2023-08-20 21:27:17+03', '2023-08-20 20:57:22+03', 'ffda6b6e-653b-87a5-7b95-924f25943f54');
INSERT INTO public.events VALUES ('4ebd4b6c-be99-652e-ab43-7fe5b15ddd6a', 'Title Notice 30', 'Description Notice 30', '2023-08-20 21:07:20+03', '2023-08-20 21:27:20+03', '2023-08-20 20:57:25+03', '143ca9e6-e441-e691-76f9-500244815245');
INSERT INTO public.events VALUES ('f90129a1-a96b-8008-3f66-486589c4a82e', 'Title Notice 32', 'Description Notice 32', '2023-08-20 21:07:21+03', '2023-08-20 21:27:21+03', '2023-08-20 20:57:26+03', 'ff59db0b-76e1-3355-99ce-203defa25983');
INSERT INTO public.events VALUES ('d45a2630-ff5b-2f42-e5c5-e23596775613', 'Title Notice 34', 'Description Notice 34', '2023-08-20 21:07:22+03', '2023-08-20 21:27:22+03', '2023-08-20 20:57:27+03', 'd9065c7d-ed80-1fa1-bb39-3d1848eee659');
INSERT INTO public.events VALUES ('4d92731c-ad64-294b-d5c7-f6effcd02f14', 'Title Notice 36', 'Description Notice 36', '2023-08-20 21:07:23+03', '2023-08-20 21:27:23+03', '2023-08-20 20:57:28+03', '4dfba4ed-22c3-e66b-6ad8-0709380b1e49');
INSERT INTO public.events VALUES ('e02f0182-6201-c095-0f09-eaa776dcdd59', 'Title Notice 38', 'Description Notice 38', '2023-08-20 21:07:24+03', '2023-08-20 21:27:24+03', '2023-08-20 20:57:29+03', '28a7aea9-0398-9f0f-f2cd-5097c7459ad3');
INSERT INTO public.events VALUES ('6dd3fb83-54a3-18cc-efe4-a018384ccbee', 'Title Notice 40', 'Description Notice 40', '2023-08-20 21:07:25+03', '2023-08-20 21:27:25+03', '2023-08-20 20:57:30+03', '51890e9f-0ece-ab9e-eb6f-4ef9de222970');
INSERT INTO public.events VALUES ('654092f2-9204-2367-6d4f-be9d4e649663', 'Title Notice 42', 'Description Notice 42', '2023-08-20 21:07:26+03', '2023-08-20 21:27:26+03', '2023-08-20 20:57:31+03', '30221d8f-8280-5e3f-1827-c8bbc75eeecb');
INSERT INTO public.events VALUES ('e50a2e0f-9caa-7681-b44e-dfafe9b41ec6', 'Title Notice 2', 'Description Notice 2', '2023-08-20 21:48:12+03', '2023-08-20 22:08:12+03', '2023-08-20 21:38:17+03', 'b72de0ec-2084-ef6d-5d9a-654ceb1c30cc');
INSERT INTO public.events VALUES ('2c82ade8-8d20-e643-b84f-a80d8862eb14', 'Title Notice 4', 'Description Notice 4', '2023-08-20 21:48:14+03', '2023-08-20 22:08:14+03', '2023-08-20 21:38:19+03', 'f90991f1-6443-3acf-75c1-6e96b6636756');
INSERT INTO public.events VALUES ('979a95a0-dd0a-594a-3e04-e3a412e8b309', 'Title Notice 6', 'Description Notice 6', '2023-08-20 21:48:16+03', '2023-08-20 22:08:16+03', '2023-08-20 21:38:21+03', '813d69c7-4073-80dd-4ed0-e9ba95b02d8b');
INSERT INTO public.events VALUES ('bd76ebbb-ee14-024a-231b-92fd388dad91', 'Title Notice 8', 'Description Notice 8', '2023-08-20 21:48:18+03', '2023-08-20 22:08:18+03', '2023-08-20 21:38:23+03', '00858459-3d71-7810-2f8c-545ecdf82c4d');
INSERT INTO public.events VALUES ('4536679d-c79a-9c8b-eb13-b0b055d12837', 'Title Notice 10', 'Description Notice 10', '2023-08-20 21:48:20+03', '2023-08-20 22:08:20+03', '2023-08-20 21:38:25+03', '132d0c3f-6818-328f-585f-f7cbb039ac9c');
INSERT INTO public.events VALUES ('98b714a4-8142-e511-42af-0b970ae65213', 'Title Notice 12', 'Description Notice 12', '2023-08-20 21:48:22+03', '2023-08-20 22:08:22+03', '2023-08-20 21:38:27+03', 'dc41b74f-b1ec-0c1a-3b21-676264b208a1');
INSERT INTO public.events VALUES ('aa0d94d2-4eae-f6d8-31c7-fc7e10a8ed75', 'Title Notice 14', 'Description Notice 14', '2023-08-20 21:48:24+03', '2023-08-20 22:08:24+03', '2023-08-20 21:38:29+03', '86c355d6-2c8a-8010-39fb-5d0830164c88');
INSERT INTO public.events VALUES ('120337dc-5f31-3d84-e809-f7e88d021e45', 'Title Notice 16', 'Description Notice 16', '2023-08-20 21:48:26+03', '2023-08-20 22:08:26+03', '2023-08-20 21:38:31+03', '88f6ee0f-f693-3f39-3602-25f868a47503');
INSERT INTO public.events VALUES ('ee0e11cd-4801-cafe-648d-58ef123be856', 'Title Notice 18', 'Description Notice 18', '2023-08-20 21:48:28+03', '2023-08-20 22:08:28+03', '2023-08-20 21:38:33+03', '98343a38-9b32-a9fd-b742-1dcada9d269b');
INSERT INTO public.events VALUES ('e06b387f-3c0b-2c1a-e168-78f51016c40b', 'Title Notice 20', 'Description Notice 20', '2023-08-20 21:48:30+03', '2023-08-20 22:08:30+03', '2023-08-20 21:38:35+03', 'ed277e89-f4c8-43c6-d70f-777329120156');
INSERT INTO public.events VALUES ('11ffd187-0088-f0cb-8932-814e88142eb3', 'Title Notice 22', 'Description Notice 22', '2023-08-20 21:48:32+03', '2023-08-20 22:08:32+03', '2023-08-20 21:38:37+03', 'a7b54489-98e8-ba58-6565-34dd29cd08ac');
INSERT INTO public.events VALUES ('9b220d61-8f77-0ef4-9627-1c66b893f787', 'Title Notice 24', 'Description Notice 24', '2023-08-20 21:48:34+03', '2023-08-20 22:08:34+03', '2023-08-20 21:38:39+03', '99865a41-5338-b9b3-6bf8-3aa8d56cdf0c');
INSERT INTO public.events VALUES ('e2049490-a624-e7ff-b6d8-2d96d6393157', 'Title Notice 26', 'Description Notice 26', '2023-08-20 21:48:36+03', '2023-08-20 22:08:36+03', '2023-08-20 21:38:41+03', 'f4ac4d18-a0bb-9688-f9a6-1f827c53555d');
INSERT INTO public.events VALUES ('e9d1c95f-b12e-7e59-8282-3c2b23045c15', 'Title Notice 28', 'Description Notice 28', '2023-08-20 21:48:38+03', '2023-08-20 22:08:38+03', '2023-08-20 21:38:43+03', 'b7fd5078-f2a1-a068-d2ed-4430ea6607ce');
INSERT INTO public.events VALUES ('58d9b607-16d7-ac06-3956-857f53dcfc26', 'Title Notice 30', 'Description Notice 30', '2023-08-20 21:48:40+03', '2023-08-20 22:08:40+03', '2023-08-20 21:38:45+03', '92212c5c-9a70-98f4-79af-7817a301997e');
INSERT INTO public.events VALUES ('f68d2c62-2dbf-ce41-04e5-a3f8c6ee7dfb', 'Title Notice 32', 'Description Notice 32', '2023-08-20 21:48:42+03', '2023-08-20 22:08:42+03', '2023-08-20 21:38:47+03', '9574132c-7781-4237-4ddb-b47b7c3227ac');
INSERT INTO public.events VALUES ('ef1b9e04-05bb-2499-957a-3601d47bffed', 'Title Notice 34', 'Description Notice 34', '2023-08-20 21:48:44+03', '2023-08-20 22:08:44+03', '2023-08-20 21:38:49+03', '9c7173cb-334f-7dbb-f080-175cd06bd19d');
INSERT INTO public.events VALUES ('8daab09b-e01b-3b2c-16fc-e334e9c48409', 'Title Notice 36', 'Description Notice 36', '2023-08-20 21:48:46+03', '2023-08-20 22:08:46+03', '2023-08-20 21:38:51+03', 'a57860a1-ffca-c47b-0137-1f9dd87fc927');
INSERT INTO public.events VALUES ('796d7bab-e4f6-785a-b4d6-abcfc116c3d1', 'Title Notice 38', 'Description Notice 38', '2023-08-20 21:48:48+03', '2023-08-20 22:08:48+03', '2023-08-20 21:38:53+03', '669c6f1b-6c1e-91c0-8021-07599ff4830a');
INSERT INTO public.events VALUES ('a7eecfde-72c4-d95b-2899-7d00483f9a13', 'Title Notice 40', 'Description Notice 40', '2023-08-20 21:48:50+03', '2023-08-20 22:08:50+03', '2023-08-20 21:38:55+03', '3a161eb7-5555-237a-b40b-9dd4f3f673e8');
INSERT INTO public.events VALUES ('0836d356-bd85-a597-7e68-5e5133efbde2', 'Title Notice 42', 'Description Notice 42', '2023-08-20 21:48:52+03', '2023-08-20 22:08:52+03', '2023-08-20 21:38:57+03', '0535fb8c-5ce1-7ccb-cc56-a7e6575ebbd0');
INSERT INTO public.events VALUES ('9365d9d0-61dd-bc8f-749c-dee171496e61', 'Title Notice 44', 'Description Notice 44', '2023-08-20 21:48:54+03', '2023-08-20 22:08:54+03', '2023-08-20 21:38:59+03', '3383821f-933a-4acc-75bb-588fdb323fa9');
INSERT INTO public.events VALUES ('685fbd75-79bd-1af3-70c1-d7edac6c0adc', 'Title Notice 46', 'Description Notice 46', '2023-08-20 21:48:56+03', '2023-08-20 22:08:56+03', '2023-08-20 21:39:01+03', '2b976432-55a3-988d-3b42-26583c7d8008');
INSERT INTO public.events VALUES ('453e3337-ad7d-631b-d241-92b9f7b67d26', 'Title Notice 48', 'Description Notice 48', '2023-08-20 21:48:58+03', '2023-08-20 22:08:58+03', '2023-08-20 21:39:03+03', '9bcb827f-745e-427f-679b-e9d5bf1226b0');
INSERT INTO public.events VALUES ('ffae01bd-e66a-2593-6461-1589ffc57683', 'Title Notice 50', 'Description Notice 50', '2023-08-20 21:49:00+03', '2023-08-20 22:09:00+03', '2023-08-20 21:39:05+03', '0250736c-cf28-0203-18e0-a180723d876c');
INSERT INTO public.events VALUES ('3f10858f-092e-58fa-e7bb-530d1fa25084', 'Title Notice 70', 'Description Notice 70', '2023-08-20 21:49:20+03', '2023-08-20 22:09:20+03', '2023-08-20 21:39:25+03', '1673492c-9b25-8ba0-9748-605536f8d6f0');
INSERT INTO public.events VALUES ('0f4eade2-3a97-3427-65b4-8f5770b70374', 'Title Notice 52', 'Description Notice 52', '2023-08-20 21:49:02+03', '2023-08-20 22:09:02+03', '2023-08-20 21:39:07+03', '3c8d4b68-eab1-cf03-dcb9-83100cb5009d');
INSERT INTO public.events VALUES ('b763ef29-a91c-35d7-c43a-4b0034cdab3b', 'Title Notice 2', 'Description Notice 2', '2023-08-20 22:11:37+03', '2023-08-20 22:31:37+03', '2023-08-20 22:01:42+03', '4d4da41c-2360-c3b8-4b60-f02aaf230de8');
INSERT INTO public.events VALUES ('b9e257d1-3a8e-9591-6015-7795ed47612d', 'Title Notice 54', 'Description Notice 54', '2023-08-20 21:49:04+03', '2023-08-20 22:09:04+03', '2023-08-20 21:39:09+03', 'ead3ec1a-ef84-9c4b-b60f-2ab2bce19dd7');
INSERT INTO public.events VALUES ('40425709-af60-005a-cdc6-40604f16ac2d', 'Title Notice 72', 'Description Notice 72', '2023-08-20 21:49:22+03', '2023-08-20 22:09:22+03', '2023-08-20 21:39:27+03', '5f365d6b-0874-89cc-9f0f-c644de7e96c4');
INSERT INTO public.events VALUES ('b4ec2c94-141b-7544-970d-2d3f648f03d7', 'Title Notice 56', 'Description Notice 56', '2023-08-20 21:49:06+03', '2023-08-20 22:09:06+03', '2023-08-20 21:39:11+03', '1cf23595-79c1-b5b2-1dee-28a97fd0bb89');
INSERT INTO public.events VALUES ('2dc5217b-0c1b-1fd4-8980-992b9a7d4d3b', 'Title Notice 58', 'Description Notice 58', '2023-08-20 21:49:08+03', '2023-08-20 22:09:08+03', '2023-08-20 21:39:13+03', 'c63d5d84-8436-329f-169a-7fb371c9ce2b');
INSERT INTO public.events VALUES ('bf11b947-7b8f-3129-84f8-c5ffd87aee7b', 'Title Notice 74', 'Description Notice 74', '2023-08-20 21:49:24+03', '2023-08-20 22:09:24+03', '2023-08-20 21:39:29+03', 'ecd1240e-5e25-f999-1948-2505d0579f62');
INSERT INTO public.events VALUES ('a1e0bbeb-620e-7862-5bee-6bdcfb515874', 'Title Notice 60', 'Description Notice 60', '2023-08-20 21:49:10+03', '2023-08-20 22:09:10+03', '2023-08-20 21:39:15+03', '1e30c6d8-29ba-71ae-9cf1-854b9adecad2');
INSERT INTO public.events VALUES ('6f7e7bad-4b68-80c6-2f11-44be82fd76f2', 'Title Notice 4', 'Description Notice 4', '2023-08-20 22:11:39+03', '2023-08-20 22:31:39+03', '2023-08-20 22:01:44+03', '621d1c8c-1469-34d9-992c-57b57265f6d3');
INSERT INTO public.events VALUES ('490e3188-ac76-3d5d-add4-dff49ab47d81', 'Title Notice 62', 'Description Notice 62', '2023-08-20 21:49:12+03', '2023-08-20 22:09:12+03', '2023-08-20 21:39:17+03', 'a8a1f782-28c3-aff6-48bb-0d3ef0e65bc1');
INSERT INTO public.events VALUES ('23c932e0-6972-8429-b5b8-33d0758801d5', 'Title Notice 76', 'Description Notice 76', '2023-08-20 21:49:26+03', '2023-08-20 22:09:26+03', '2023-08-20 21:39:31+03', '2ab35c11-3607-cc4f-e8e8-91ee7d590f8c');
INSERT INTO public.events VALUES ('debd5ba8-ad96-17a8-1a59-ad5e769b7eb0', 'Title Notice 64', 'Description Notice 64', '2023-08-20 21:49:14+03', '2023-08-20 22:09:14+03', '2023-08-20 21:39:19+03', '64de7695-4b55-fa87-f8c1-74d223d9b34d');
INSERT INTO public.events VALUES ('b0db4739-a5f4-0ab5-b732-cc85edd29278', 'Title Notice 66', 'Description Notice 66', '2023-08-20 21:49:16+03', '2023-08-20 22:09:16+03', '2023-08-20 21:39:21+03', 'ddb5d10d-e994-18e8-4c05-8b2ca2573237');
INSERT INTO public.events VALUES ('d3a89cab-1e38-2afe-f9db-73dc691d33bc', 'Title Notice 78', 'Description Notice 78', '2023-08-20 21:49:28+03', '2023-08-20 22:09:28+03', '2023-08-20 21:39:33+03', '2661b5a0-7091-82ad-b988-e75c0c8af724');
INSERT INTO public.events VALUES ('96726d21-2914-a505-4171-28399478588a', 'Title Notice 68', 'Description Notice 68', '2023-08-20 21:49:18+03', '2023-08-20 22:09:18+03', '2023-08-20 21:39:23+03', '17bf8e81-02a9-779b-9512-4c844718dd4e');
INSERT INTO public.events VALUES ('c857f4c4-338b-1ce6-e098-a481b6d8fbee', 'Title Notice 6', 'Description Notice 6', '2023-08-20 22:11:41+03', '2023-08-20 22:31:41+03', '2023-08-20 22:01:46+03', 'b12fecb0-236a-456d-3849-2ec42e4b7274');
INSERT INTO public.events VALUES ('6c5660bf-0925-04d6-606b-1cc0917929f4', 'Title Notice 80', 'Description Notice 80', '2023-08-20 21:49:30+03', '2023-08-20 22:09:30+03', '2023-08-20 21:39:35+03', '142af404-7abe-959e-ea3a-18f5c43090bc');
INSERT INTO public.events VALUES ('ef087c81-04d1-7c03-d626-10b517cc8654', 'Title Notice 82', 'Description Notice 82', '2023-08-20 21:49:32+03', '2023-08-20 22:09:32+03', '2023-08-20 21:39:37+03', 'bcd510bb-c6b6-8b93-546a-95c8875b9934');
INSERT INTO public.events VALUES ('60cbeb4e-5c78-a978-1c51-757bac567671', 'Title Notice 8', 'Description Notice 8', '2023-08-20 22:11:43+03', '2023-08-20 22:31:43+03', '2023-08-20 22:01:48+03', '20733943-2cb6-80dc-0994-0c6d61068dd1');
INSERT INTO public.events VALUES ('7c1f0c4d-9c35-c222-342c-756752c2b67a', 'Title Notice 84', 'Description Notice 84', '2023-08-20 21:49:34+03', '2023-08-20 22:09:34+03', '2023-08-20 21:39:39+03', 'b3b18829-7b62-9d48-3f25-187cd7d9993f');
INSERT INTO public.events VALUES ('a1b921a1-806d-6969-c047-12d46ad8996f', 'Title Notice 86', 'Description Notice 86', '2023-08-20 21:49:36+03', '2023-08-20 22:09:36+03', '2023-08-20 21:39:41+03', '9d40e8df-3ca6-911c-1a58-7089bfdc4c39');
INSERT INTO public.events VALUES ('65be946f-8f3f-7da8-e73c-639218e4c130', 'Title Notice 10', 'Description Notice 10', '2023-08-20 22:11:45+03', '2023-08-20 22:31:45+03', '2023-08-20 22:01:50+03', 'f4cf18e9-1956-277b-559c-aca022bbee6a');
INSERT INTO public.events VALUES ('f7ec9079-e34d-e4f0-c18a-965e2ca38b90', 'Title Notice 88', 'Description Notice 88', '2023-08-20 21:49:38+03', '2023-08-20 22:09:38+03', '2023-08-20 21:39:43+03', '4bad344f-7a39-72cd-345e-3cf0ea1c5a17');
INSERT INTO public.events VALUES ('7ed0082a-1f0c-eb1c-8dee-39af64c1ba0d', 'Title Notice 90', 'Description Notice 90', '2023-08-20 21:49:40+03', '2023-08-20 22:09:40+03', '2023-08-20 21:39:45+03', '036fe37b-56ce-2609-fb69-97862999d444');
INSERT INTO public.events VALUES ('5a8b5f09-fa72-59a3-12e3-b1513bca54e2', 'Title Notice 92', 'Description Notice 92', '2023-08-20 21:49:42+03', '2023-08-20 22:09:42+03', '2023-08-20 21:39:47+03', '67ff9b0b-688b-4155-71d0-0b66b0535ba4');
INSERT INTO public.events VALUES ('d2337537-7c2f-b7fa-e44d-a298baf5c97a', 'Title Notice 94', 'Description Notice 94', '2023-08-20 21:49:44+03', '2023-08-20 22:09:44+03', '2023-08-20 21:39:49+03', '0f1612a6-b57f-7f83-f71e-094979f9943e');
INSERT INTO public.events VALUES ('ad25b490-af89-80a9-4080-25f439eaf5b3', 'Title Notice 96', 'Description Notice 96', '2023-08-20 21:49:46+03', '2023-08-20 22:09:46+03', '2023-08-20 21:39:51+03', '90e4bb14-5202-d7d7-d174-4dd25944e2a0');
INSERT INTO public.events VALUES ('6423607a-7ee5-f0b3-bb9c-aaf0361438c1', 'Title Notice 98', 'Description Notice 98', '2023-08-20 21:49:48+03', '2023-08-20 22:09:48+03', '2023-08-20 21:39:53+03', '72e71bb0-141a-faf4-3995-d28ef7fdee90');
INSERT INTO public.events VALUES ('ec6fccf2-9c3e-975b-e983-79af5f90d3a0', 'Title Notice 100', 'Description Notice 100', '2023-08-20 21:49:50+03', '2023-08-20 22:09:50+03', '2023-08-20 21:39:55+03', '458a4f9e-44a0-14a6-c767-362c67be2138');
INSERT INTO public.events VALUES ('66bdcf3b-6e03-c413-9202-0f313619547e', 'Title Notice 102', 'Description Notice 102', '2023-08-20 21:49:52+03', '2023-08-20 22:09:52+03', '2023-08-20 21:39:57+03', '35cb0a71-4bf6-e786-3fbb-a139ac79107b');
INSERT INTO public.events VALUES ('f46fb33c-c567-de8e-a84e-6adfdc62e088', 'Title Notice 104', 'Description Notice 104', '2023-08-20 21:49:54+03', '2023-08-20 22:09:54+03', '2023-08-20 21:39:59+03', '51c018cc-b6c0-de9e-934d-b04d72836511');
INSERT INTO public.events VALUES ('94c40f1b-86bd-11c7-7d7a-c9671782ef0e', 'Title Notice 106', 'Description Notice 106', '2023-08-20 21:49:56+03', '2023-08-20 22:09:56+03', '2023-08-20 21:40:01+03', '9496c8a6-0bee-ab11-f08b-ed54f5813b9b');
INSERT INTO public.events VALUES ('58fc2d7e-cae4-b599-2748-e10f1d716b79', 'Title Notice 108', 'Description Notice 108', '2023-08-20 21:49:58+03', '2023-08-20 22:09:58+03', '2023-08-20 21:40:03+03', '12994f7a-e5b3-af89-624b-290d08a7cb4f');
INSERT INTO public.events VALUES ('916728a3-f553-164f-97e0-967a99c5e439', 'Title Notice 110', 'Description Notice 110', '2023-08-20 21:50:00+03', '2023-08-20 22:10:00+03', '2023-08-20 21:40:05+03', '1671108c-ccd8-eb60-4d3c-01c1c90df140');
INSERT INTO public.events VALUES ('f93434aa-5dcb-f083-c2f5-bda319243295', 'Title Notice 112', 'Description Notice 112', '2023-08-20 21:50:02+03', '2023-08-20 22:10:02+03', '2023-08-20 21:40:07+03', '35cee0a7-d84a-b4fd-d8ee-ba71539188f6');
INSERT INTO public.events VALUES ('b9cc6efa-2f6b-999c-6b0b-4d5af588e794', 'Title Notice 114', 'Description Notice 114', '2023-08-20 21:50:04+03', '2023-08-20 22:10:04+03', '2023-08-20 21:40:09+03', '611aed98-ecf4-4db3-9772-6245544a4e2c');
INSERT INTO public.events VALUES ('4d1047a1-0b02-c74a-bcac-3b7fd326f5b2', 'Title Notice 116', 'Description Notice 116', '2023-08-20 21:50:06+03', '2023-08-20 22:10:06+03', '2023-08-20 21:40:11+03', '922b2390-cb7d-6564-c960-257500b8bd4b');
INSERT INTO public.events VALUES ('376670d1-5730-14cb-3b5a-db598c4e7aba', 'Title Notice 118', 'Description Notice 118', '2023-08-20 21:50:08+03', '2023-08-20 22:10:08+03', '2023-08-20 21:40:13+03', '8f12e7e7-7c3e-9d4a-3d99-69626bd36514');
INSERT INTO public.events VALUES ('b58340d5-88d3-11d6-816b-7c435b0f4a7b', 'Title Notice 120', 'Description Notice 120', '2023-08-20 21:50:10+03', '2023-08-20 22:10:10+03', '2023-08-20 21:40:15+03', 'f67d74f7-6452-fd6e-dc37-65518dac556e');
INSERT INTO public.events VALUES ('ff6e8540-17f0-e595-b2be-3b9052657d92', 'Title Notice 122', 'Description Notice 122', '2023-08-20 21:50:12+03', '2023-08-20 22:10:12+03', '2023-08-20 21:40:17+03', '9a8f2287-a8d2-db9e-5142-79908e38ed6b');
INSERT INTO public.events VALUES ('1502fa38-984c-d37b-2daa-0f5e32b79940', 'Title Notice 124', 'Description Notice 124', '2023-08-20 21:50:14+03', '2023-08-20 22:10:14+03', '2023-08-20 21:40:19+03', '5e490801-8433-ea5c-ac92-d0822017aa6d');
INSERT INTO public.events VALUES ('0ab9f49e-af6c-2d48-b3fe-80794065f7cf', 'Title Notice 126', 'Description Notice 126', '2023-08-20 21:50:16+03', '2023-08-20 22:10:16+03', '2023-08-20 21:40:21+03', 'bc97df86-36d1-a7c0-c2e5-842f8d5ec049');
INSERT INTO public.events VALUES ('0dcb6d6a-6db1-3d11-1a87-bf84ce427a69', 'Title Notice 128', 'Description Notice 128', '2023-08-20 21:50:18+03', '2023-08-20 22:10:18+03', '2023-08-20 21:40:23+03', 'a4920906-2978-0aba-03e2-d2c6f4e74807');
INSERT INTO public.events VALUES ('4005e723-f200-d853-f6f4-d3cbae606bfc', 'Title Notice 12', 'Description Notice 12', '2023-08-20 22:11:47+03', '2023-08-20 22:31:47+03', '2023-08-20 22:01:52+03', 'ab6d54de-4155-4c18-8e4d-cf6f3a2effdb');
INSERT INTO public.events VALUES ('4adad789-6aca-9f91-a383-2ec7b12641ea', 'Title Notice 130', 'Description Notice 130', '2023-08-20 21:50:20+03', '2023-08-20 22:10:20+03', '2023-08-20 21:40:25+03', '201e29ab-18d7-78cf-1789-fe673157564e');
INSERT INTO public.events VALUES ('e88bbb13-6a5a-ee51-d334-fa1b5bbeea94', 'Title Notice 132', 'Description Notice 132', '2023-08-20 21:50:22+03', '2023-08-20 22:10:22+03', '2023-08-20 21:40:27+03', '8016ecfe-8201-0609-2d24-9a5a2d49f964');
INSERT INTO public.events VALUES ('3830ff0a-7059-edd9-38f8-c1742a61f127', 'Title Notice 14', 'Description Notice 14', '2023-08-20 22:11:49+03', '2023-08-20 22:31:49+03', '2023-08-20 22:01:54+03', 'cc89fc0c-7ee2-f7aa-4306-04edaf3120df');
INSERT INTO public.events VALUES ('668ef80c-e306-0176-1ba4-2376bfbdf600', 'Title Notice 134', 'Description Notice 134', '2023-08-20 21:50:24+03', '2023-08-20 22:10:24+03', '2023-08-20 21:40:29+03', 'e6341b15-5bf8-f221-14fe-4c8a49cb247d');
INSERT INTO public.events VALUES ('92f5df37-5d67-dce1-0250-e014fb38d6da', 'Title Notice 2', 'Description Notice 2', '2023-08-20 21:59:02+03', '2023-08-20 22:19:02+03', '2023-08-20 21:49:07+03', 'c19c4f02-5817-332f-6144-bfa24ad303f4');
INSERT INTO public.events VALUES ('645ea984-ed7a-31da-8a02-8396b60396cd', 'Title Notice 16', 'Description Notice 16', '2023-08-20 22:11:51+03', '2023-08-20 22:31:51+03', '2023-08-20 22:01:56+03', 'e0c4b50c-3677-2cf8-4926-17d3dbef47bb');
INSERT INTO public.events VALUES ('6f62f340-b1fb-566a-d19c-9de6204fec05', 'Title Notice 4', 'Description Notice 4', '2023-08-20 21:59:04+03', '2023-08-20 22:19:04+03', '2023-08-20 21:49:09+03', '49c5631e-0496-2f03-9f3c-fd8f053946ae');
INSERT INTO public.events VALUES ('3728d15d-966a-e80b-b46a-da7c69c2ea41', 'Title Notice 6', 'Description Notice 6', '2023-08-20 21:59:06+03', '2023-08-20 22:19:06+03', '2023-08-20 21:49:11+03', '014c5e33-3f0e-1250-7ca3-08cfe4fe4a73');
INSERT INTO public.events VALUES ('29dd00d0-dcd2-a96f-7046-1b9952041de3', 'Title Notice 8', 'Description Notice 8', '2023-08-20 21:59:08+03', '2023-08-20 22:19:08+03', '2023-08-20 21:49:13+03', 'c5f9c7a4-3e69-07c0-583f-35a57421a134');
INSERT INTO public.events VALUES ('3758284a-2ae0-1dd1-8ffb-64d7a3dab35c', 'Title Notice 2', 'Description Notice 2', '2023-08-20 21:59:32+03', '2023-08-20 22:19:32+03', '2023-08-20 21:49:37+03', 'f142c03e-0cf7-0e92-57bb-f7e0a95113ff');
INSERT INTO public.events VALUES ('400542c4-9f5e-979e-3197-508fadea15ea', 'Title Notice 4', 'Description Notice 4', '2023-08-20 21:59:34+03', '2023-08-20 22:19:34+03', '2023-08-20 21:49:39+03', 'c833f989-f542-7279-7e6e-d73b2456000c');
INSERT INTO public.events VALUES ('4c32f7ba-32e4-ba4a-f7d4-7130396eb4b7', 'Title Notice 6', 'Description Notice 6', '2023-08-20 21:59:36+03', '2023-08-20 22:19:36+03', '2023-08-20 21:49:41+03', '09fd38d5-b260-c406-5594-ca145700d2bd');
INSERT INTO public.events VALUES ('e072ad5d-8f5e-da7c-6b11-6a5e9c41b16e', 'Title Notice 8', 'Description Notice 8', '2023-08-20 21:59:38+03', '2023-08-20 22:19:38+03', '2023-08-20 21:49:43+03', 'd8ca3f1c-a874-3738-0568-5c1e9d1bea2e');
INSERT INTO public.events VALUES ('bb6619f2-0477-3b53-414e-9fb17131bcb4', 'Title Notice 10', 'Description Notice 10', '2023-08-20 21:59:40+03', '2023-08-20 22:19:40+03', '2023-08-20 21:49:45+03', '7313f648-5387-5861-ed26-33bb08707cca');
INSERT INTO public.events VALUES ('8550a9a3-fd6b-464a-41b2-843d85cff61b', 'Title Notice 12', 'Description Notice 12', '2023-08-20 21:59:42+03', '2023-08-20 22:19:42+03', '2023-08-20 21:49:47+03', 'e0e1affe-1cc9-a331-dc39-88b1569a3990');
INSERT INTO public.events VALUES ('e87ea900-31d2-d823-b1ab-1e4df342f01b', 'Title Notice 14', 'Description Notice 14', '2023-08-20 21:59:44+03', '2023-08-20 22:19:44+03', '2023-08-20 21:49:49+03', '29edf610-d646-f2a1-fb7f-565fc94d8c36');
INSERT INTO public.events VALUES ('293faa88-b0a9-490c-a9b1-6446ea897843', 'Title Notice 16', 'Description Notice 16', '2023-08-20 21:59:46+03', '2023-08-20 22:19:46+03', '2023-08-20 21:49:51+03', '2033210b-33f1-cb40-840c-760e12fa6c8a');
INSERT INTO public.events VALUES ('b267a3ce-6fc7-3203-273c-e89061cacdab', 'Title Notice 18', 'Description Notice 18', '2023-08-20 21:59:48+03', '2023-08-20 22:19:48+03', '2023-08-20 21:49:53+03', '4462d47d-d086-b10a-4e2a-6c9da9ec7dfd');
INSERT INTO public.events VALUES ('930751ea-80ca-8add-ac17-485fe854ecbd', 'Title Notice 20', 'Description Notice 20', '2023-08-20 21:59:50+03', '2023-08-20 22:19:50+03', '2023-08-20 21:49:55+03', 'e261cc1b-0627-a602-706a-6f05424461b4');
INSERT INTO public.events VALUES ('b6ffe12e-1bc1-233b-81b4-ea60510d3116', 'Title Notice 22', 'Description Notice 22', '2023-08-20 21:59:52+03', '2023-08-20 22:19:52+03', '2023-08-20 21:49:57+03', '26cd8eb9-11f9-db71-e006-5ad8301affca');
INSERT INTO public.events VALUES ('42225584-274f-151c-c454-838d90e03660', 'Title Notice 24', 'Description Notice 24', '2023-08-20 21:59:54+03', '2023-08-20 22:19:54+03', '2023-08-20 21:49:59+03', '05459140-1417-851c-5edd-2e3a169c095d');
INSERT INTO public.events VALUES ('e45bf4ad-1940-6780-f524-4a8e2c16b4b4', 'Title Notice 26', 'Description Notice 26', '2023-08-20 21:59:56+03', '2023-08-20 22:19:56+03', '2023-08-20 21:50:01+03', '64503c20-7831-46ba-cf8a-f71165bdb216');
INSERT INTO public.events VALUES ('337c8d85-94b9-9ee0-d661-39efaec6bc14', 'Title Notice 28', 'Description Notice 28', '2023-08-20 21:59:58+03', '2023-08-20 22:19:58+03', '2023-08-20 21:50:03+03', '5bd7c0d0-775a-f95e-1285-d26edcd783f3');
INSERT INTO public.events VALUES ('6b6da10a-43a6-3db4-453d-ac1f17858349', 'Title Notice 30', 'Description Notice 30', '2023-08-20 22:00:00+03', '2023-08-20 22:20:00+03', '2023-08-20 21:50:05+03', '865d4728-040c-b011-1d68-a382c8711647');
INSERT INTO public.events VALUES ('8f7ca4c7-8464-96b8-7b70-e707f18a174d', 'Title Notice 32', 'Description Notice 32', '2023-08-20 22:00:02+03', '2023-08-20 22:20:02+03', '2023-08-20 21:50:07+03', '018d2b85-d70d-412f-d8a5-9ab4b859b27a');
INSERT INTO public.events VALUES ('82c0096a-6e11-395a-3436-9000c9101155', 'Title Notice 34', 'Description Notice 34', '2023-08-20 22:00:04+03', '2023-08-20 22:20:04+03', '2023-08-20 21:50:09+03', 'c374d5b3-fc9c-0f15-9a62-447ed561ff2b');
INSERT INTO public.events VALUES ('6dbe6fc2-26e8-54bb-bf13-81d41a8e425b', 'Title Notice 36', 'Description Notice 36', '2023-08-20 22:00:06+03', '2023-08-20 22:20:06+03', '2023-08-20 21:50:11+03', '4cfc208c-66de-174f-bdde-81b8da38f8b7');
INSERT INTO public.events VALUES ('00b10fba-7851-ef22-cecc-b6b93bf4cf5a', 'Title Notice 38', 'Description Notice 38', '2023-08-20 22:00:08+03', '2023-08-20 22:20:08+03', '2023-08-20 21:50:13+03', '883cb2ab-ad82-5e01-810b-13bbf665f205');
INSERT INTO public.events VALUES ('ed5bd53a-39fc-5b51-c661-305c81c6703f', 'Title Notice 40', 'Description Notice 40', '2023-08-20 22:00:10+03', '2023-08-20 22:20:10+03', '2023-08-20 21:50:15+03', '0d4a40c2-7afb-bae6-ad95-364f59322035');
INSERT INTO public.events VALUES ('054eb355-e47c-7ece-05de-5bf31fac2aba', 'Title Notice 42', 'Description Notice 42', '2023-08-20 22:00:12+03', '2023-08-20 22:20:12+03', '2023-08-20 21:50:17+03', '26276461-5de8-fad4-052e-deeab57cb2a8');
INSERT INTO public.events VALUES ('f5d5c322-93dd-ba18-c4db-a5b9ca2dcde3', 'Title Notice 44', 'Description Notice 44', '2023-08-20 22:00:14+03', '2023-08-20 22:20:14+03', '2023-08-20 21:50:19+03', 'd4364ac5-ca7f-3979-c7fe-33af95cb6b75');
INSERT INTO public.events VALUES ('eac1c779-3635-9c58-4541-8f9d334c40a0', 'Title Notice 46', 'Description Notice 46', '2023-08-20 22:00:16+03', '2023-08-20 22:20:16+03', '2023-08-20 21:50:21+03', 'f61ff6cf-02e0-e9a8-3c62-4d94aff2491a');
INSERT INTO public.events VALUES ('0ed30407-817b-bd62-ea77-942ae85b4d26', 'Title Notice 48', 'Description Notice 48', '2023-08-20 22:00:18+03', '2023-08-20 22:20:18+03', '2023-08-20 21:50:23+03', 'd5d8017f-f9c3-d25a-faa7-9d4c4ecaca83');
INSERT INTO public.events VALUES ('a80a329b-cfc3-b376-8fef-bbce25aab1b3', 'Title Notice 50', 'Description Notice 50', '2023-08-20 22:00:20+03', '2023-08-20 22:20:20+03', '2023-08-20 21:50:25+03', '5985ebfe-48f7-e093-fc08-bf10dc8016ba');
INSERT INTO public.events VALUES ('c77a88c9-5bb1-3aa7-d6f0-262e8cda1c6c', 'Title Notice 52', 'Description Notice 52', '2023-08-20 22:00:22+03', '2023-08-20 22:20:22+03', '2023-08-20 21:50:27+03', 'f7553449-6e12-c981-819b-51fb3d68752c');
INSERT INTO public.events VALUES ('f4bb62f9-1f65-a153-3fb7-0ad48172c037', 'Title Notice 54', 'Description Notice 54', '2023-08-20 22:00:24+03', '2023-08-20 22:20:24+03', '2023-08-20 21:50:29+03', '5fb068c2-bd16-07bc-f63a-f7e12e0bc38d');
INSERT INTO public.events VALUES ('ceb76ae0-48b0-c409-fcaa-cfeb93cb6862', 'Title Notice 56', 'Description Notice 56', '2023-08-20 22:00:26+03', '2023-08-20 22:20:26+03', '2023-08-20 21:50:31+03', '3e0077c6-f4fa-8ec3-25bd-1dd142bf83f8');
INSERT INTO public.events VALUES ('046a5be5-b509-a523-aa9f-7e201f28042c', 'Title Notice 58', 'Description Notice 58', '2023-08-20 22:00:28+03', '2023-08-20 22:20:28+03', '2023-08-20 21:50:33+03', 'c59e32c2-59fa-8c79-4add-708939117892');
INSERT INTO public.events VALUES ('35c3a0d5-add2-3851-e7ef-e0155353c99e', 'Title Notice 60', 'Description Notice 60', '2023-08-20 22:00:30+03', '2023-08-20 22:20:30+03', '2023-08-20 21:50:35+03', 'a51a5c61-49e8-05ab-1add-d146c2d817d5');
INSERT INTO public.events VALUES ('c843db46-2372-974c-8899-6258567476d2', 'Title Notice 62', 'Description Notice 62', '2023-08-20 22:00:32+03', '2023-08-20 22:20:32+03', '2023-08-20 21:50:37+03', '65e2a4d0-2c91-7a03-ac22-08a81919ad08');
INSERT INTO public.events VALUES ('e3a191bf-df05-bbf0-adb2-286df839c1e4', 'Title Notice 64', 'Description Notice 64', '2023-08-20 22:00:34+03', '2023-08-20 22:20:34+03', '2023-08-20 21:50:39+03', '7f1b35d8-3318-8a25-1636-ff32384790b8');
INSERT INTO public.events VALUES ('3e47871b-f44c-fa0d-9b39-5e499b6bb0b5', 'Title Notice 66', 'Description Notice 66', '2023-08-20 22:00:36+03', '2023-08-20 22:20:36+03', '2023-08-20 21:50:41+03', 'bed295a6-5850-8b90-b046-7405733611ba');
INSERT INTO public.events VALUES ('becde445-5495-3740-06f0-ad12ea4b06b4', 'Title Notice 68', 'Description Notice 68', '2023-08-20 22:00:38+03', '2023-08-20 22:20:38+03', '2023-08-20 21:50:43+03', 'ca592cf3-d1b7-3a86-6dee-47e15aaf83bb');
INSERT INTO public.events VALUES ('af2310eb-026d-1679-43f8-65e61535e717', 'Title Notice 70', 'Description Notice 70', '2023-08-20 22:00:40+03', '2023-08-20 22:20:40+03', '2023-08-20 21:50:45+03', '6475d304-f410-3b49-f9a0-6aa8f9f18971');
INSERT INTO public.events VALUES ('f2ac2e33-0678-a43a-16f8-8da08716e44f', 'Title Notice 18', 'Description Notice 18', '2023-08-20 22:11:53+03', '2023-08-20 22:31:53+03', '2023-08-20 22:01:58+03', '8335fa6c-e7d3-9f79-d030-06fa59e775e5');
INSERT INTO public.events VALUES ('720ff1be-43a2-fa3c-3663-6e0e779e776b', 'Title Notice 72', 'Description Notice 72', '2023-08-20 22:00:42+03', '2023-08-20 22:20:42+03', '2023-08-20 21:50:47+03', 'd8606221-ec85-1f83-413d-3115d9841f33');
INSERT INTO public.events VALUES ('cd9ceb7f-528e-726d-07e4-9398bbfec9e1', 'Title Notice 74', 'Description Notice 74', '2023-08-20 22:00:44+03', '2023-08-20 22:20:44+03', '2023-08-20 21:50:49+03', 'ba48369c-a3e9-bbd2-6394-1df736b52a34');
INSERT INTO public.events VALUES ('94ba3fd6-33c4-4926-f01a-732cd4424408', 'Title Notice 20', 'Description Notice 20', '2023-08-20 22:11:55+03', '2023-08-20 22:31:55+03', '2023-08-20 22:02:00+03', '8230dd57-587a-ea17-06a8-df8f7424864e');
INSERT INTO public.events VALUES ('4bf7987e-a8a9-be5d-cfa0-9e64a1bf0f80', 'Title Notice 76', 'Description Notice 76', '2023-08-20 22:00:46+03', '2023-08-20 22:20:46+03', '2023-08-20 21:50:51+03', 'ca6983a7-80f7-a8dc-243d-71fa14904234');
INSERT INTO public.events VALUES ('9d8e3e83-2c66-8519-23af-e60f11c478ac', 'Title Notice 78', 'Description Notice 78', '2023-08-20 22:00:48+03', '2023-08-20 22:20:48+03', '2023-08-20 21:50:53+03', '58b19958-a45b-ab17-6553-5ad42bfc2133');
INSERT INTO public.events VALUES ('9570b8fc-e055-8c97-261f-e67f21a6d83b', 'Title Notice 22', 'Description Notice 22', '2023-08-20 22:11:57+03', '2023-08-20 22:31:57+03', '2023-08-20 22:02:02+03', '933f2d44-66fb-75ab-a4f2-9da4af9446c4');
INSERT INTO public.events VALUES ('b1da0c3f-9c0e-b860-f0b5-a9ae3906fc70', 'Title Notice 80', 'Description Notice 80', '2023-08-20 22:00:50+03', '2023-08-20 22:20:50+03', '2023-08-20 21:50:55+03', 'dc013e62-09bf-7707-02f5-f767b7a04621');
INSERT INTO public.events VALUES ('87e98255-4350-5991-c653-d74e37bfda66', 'Title Notice 82', 'Description Notice 82', '2023-08-20 22:00:52+03', '2023-08-20 22:20:52+03', '2023-08-20 21:50:57+03', 'e5450821-c9e8-e9e4-3b9c-53c8c0c13213');
INSERT INTO public.events VALUES ('83845569-0dd2-d6bc-715f-bbf606ae2a8b', 'Title Notice 24', 'Description Notice 24', '2023-08-20 22:11:59+03', '2023-08-20 22:31:59+03', '2023-08-20 22:02:04+03', 'd2120832-9de5-ddb9-f695-b2968400ff6d');
INSERT INTO public.events VALUES ('bd6b2331-f8f6-daae-73c4-a06a3d50020a', 'Title Notice 84', 'Description Notice 84', '2023-08-20 22:00:54+03', '2023-08-20 22:20:54+03', '2023-08-20 21:50:59+03', 'e801dfac-53c6-89fd-5112-1e0ed70c599f');
INSERT INTO public.events VALUES ('faa428e5-96bf-cb1b-17fe-d9477ae8e573', 'Title Notice 86', 'Description Notice 86', '2023-08-20 22:00:56+03', '2023-08-20 22:20:56+03', '2023-08-20 21:51:01+03', 'e22387b9-aef4-90e4-8240-fa1e1d05975b');
INSERT INTO public.events VALUES ('629b1d7a-ba1c-8e6b-2967-3f7482ab19b2', 'Title Notice 26', 'Description Notice 26', '2023-08-20 22:12:01+03', '2023-08-20 22:32:01+03', '2023-08-20 22:02:06+03', '71c13aaa-dd18-8355-3808-725449ffa93d');
INSERT INTO public.events VALUES ('b8bcde8e-c2f5-7010-2482-a6ac42c949d8', 'Title Notice 88', 'Description Notice 88', '2023-08-20 22:00:58+03', '2023-08-20 22:20:58+03', '2023-08-20 21:51:03+03', 'c883c143-2a79-1163-3b37-75ea5d541025');
INSERT INTO public.events VALUES ('092994c3-d895-6cd3-2e59-9d773a77602d', 'Title Notice 90', 'Description Notice 90', '2023-08-20 22:01:00+03', '2023-08-20 22:21:00+03', '2023-08-20 21:51:05+03', '556c254f-0480-8f33-501b-3ffa6667c619');
INSERT INTO public.events VALUES ('c155893f-f018-6266-1f3e-6e168000bf35', 'Title Notice 28', 'Description Notice 28', '2023-08-20 22:12:03+03', '2023-08-20 22:32:03+03', '2023-08-20 22:02:08+03', '5b20bab5-d40c-838f-f6f3-5564ed69f0fe');
INSERT INTO public.events VALUES ('195c3bff-15d6-edfb-6297-922e9598249b', 'Title Notice 92', 'Description Notice 92', '2023-08-20 22:01:02+03', '2023-08-20 22:21:02+03', '2023-08-20 21:51:07+03', 'ae4e1c82-e709-05eb-f485-8dcd51957ff5');
INSERT INTO public.events VALUES ('44a2aae4-0009-17cc-b3a6-a028b7d808d2', 'Title Notice 94', 'Description Notice 94', '2023-08-20 22:01:04+03', '2023-08-20 22:21:04+03', '2023-08-20 21:51:09+03', 'da35764f-1688-bc55-6355-7dea4c2d964e');
INSERT INTO public.events VALUES ('69be497b-0579-4469-7c4a-77ff6eb4c5b1', 'Title Notice 30', 'Description Notice 30', '2023-08-20 22:12:05+03', '2023-08-20 22:32:05+03', '2023-08-20 22:02:10+03', 'cf0e91d6-4793-4a8f-4158-2f72940c5b9c');
INSERT INTO public.events VALUES ('531c1047-40f7-c1fa-9ad5-f581fc663a3d', 'Title Notice 96', 'Description Notice 96', '2023-08-20 22:01:06+03', '2023-08-20 22:21:06+03', '2023-08-20 21:51:11+03', '472ee4f0-0750-f075-284b-2b5e6c048a47');
INSERT INTO public.events VALUES ('a3817254-36b3-3766-37dc-e3fcbd73e7c3', 'Title Notice 98', 'Description Notice 98', '2023-08-20 22:01:08+03', '2023-08-20 22:21:08+03', '2023-08-20 21:51:13+03', '7df320b7-c02e-a13f-7abd-f352c6726f1c');
INSERT INTO public.events VALUES ('5bedd4d0-9587-0439-ec5b-38844d99e163', 'Title Notice 32', 'Description Notice 32', '2023-08-20 22:12:07+03', '2023-08-20 22:32:07+03', '2023-08-20 22:02:12+03', '1ed9400b-1a6c-8678-4192-0cf0cbddbc8d');
INSERT INTO public.events VALUES ('65ff27aa-f7c2-3113-40a3-97da50feffe0', 'Title Notice 100', 'Description Notice 100', '2023-08-20 22:01:10+03', '2023-08-20 22:21:10+03', '2023-08-20 21:51:15+03', '27e5e1bb-71d2-9d31-ff5b-ee4923601cd3');
INSERT INTO public.events VALUES ('b3482129-c539-4973-e7a0-9b142ec5b049', 'Title Notice 102', 'Description Notice 102', '2023-08-20 22:01:12+03', '2023-08-20 22:21:12+03', '2023-08-20 21:51:17+03', '94e241b5-e8ee-f101-25e1-b301e6cf4451');
INSERT INTO public.events VALUES ('18a05c47-3382-a02e-15ad-957a5dccdc9a', 'Title Notice 34', 'Description Notice 34', '2023-08-20 22:12:09+03', '2023-08-20 22:32:09+03', '2023-08-20 22:02:14+03', 'ce37e90e-6e18-8be6-21ec-c6593b7f798d');
INSERT INTO public.events VALUES ('e898b71e-13d4-8815-77db-1778cf3c203e', 'Title Notice 104', 'Description Notice 104', '2023-08-20 22:01:14+03', '2023-08-20 22:21:14+03', '2023-08-20 21:51:19+03', '442a4ed5-a383-352c-c8f7-b0ab6c7a836b');
INSERT INTO public.events VALUES ('b5c9db2d-d5b2-ad0a-03bf-baf2d5c47b80', 'Title Notice 106', 'Description Notice 106', '2023-08-20 22:01:16+03', '2023-08-20 22:21:16+03', '2023-08-20 21:51:21+03', 'aa92e4be-43e4-aafc-ac5f-121ce1b68e65');
INSERT INTO public.events VALUES ('3219d08e-844f-0b7d-67c4-b7dd3b29d645', 'Title Notice 36', 'Description Notice 36', '2023-08-20 22:12:11+03', '2023-08-20 22:32:11+03', '2023-08-20 22:02:16+03', '3dcf0fd5-d68f-75d2-7306-ac51af22035b');
INSERT INTO public.events VALUES ('62167f2b-eabe-8de5-55af-9b635f2b16f9', 'Title Notice 108', 'Description Notice 108', '2023-08-20 22:01:18+03', '2023-08-20 22:21:18+03', '2023-08-20 21:51:23+03', '8d70167f-a912-a89b-eedf-806d6b43e1f1');
INSERT INTO public.events VALUES ('19c39bce-b9e1-0e75-ddcb-de625a386c07', 'Title Notice 110', 'Description Notice 110', '2023-08-20 22:01:20+03', '2023-08-20 22:21:20+03', '2023-08-20 21:51:25+03', 'abbc5e36-ffc9-2130-ede1-8461bb2ce711');
INSERT INTO public.events VALUES ('7307bf53-68a2-7559-acdc-b83d5078a208', 'Title Notice 38', 'Description Notice 38', '2023-08-20 22:12:13+03', '2023-08-20 22:32:13+03', '2023-08-20 22:02:18+03', '56bbece3-dfa7-9578-f232-c90e2791591a');
INSERT INTO public.events VALUES ('77863dab-bb63-dc62-9c28-fba36007163a', 'Title Notice 112', 'Description Notice 112', '2023-08-20 22:01:22+03', '2023-08-20 22:21:22+03', '2023-08-20 21:51:27+03', '763d521d-efed-dd3f-f013-1b36f15dcd84');
INSERT INTO public.events VALUES ('c783bdb6-d07e-e226-4f8b-242ec83db0b7', 'Title Notice 114', 'Description Notice 114', '2023-08-20 22:01:24+03', '2023-08-20 22:21:24+03', '2023-08-20 21:51:29+03', 'd9a4154b-a419-ed03-cf27-754bbe4327f4');
INSERT INTO public.events VALUES ('4ae31838-ac41-0180-adeb-97da279f3553', 'Title Notice 40', 'Description Notice 40', '2023-08-20 22:12:15+03', '2023-08-20 22:32:15+03', '2023-08-20 22:02:20+03', 'eda57908-f01c-4f9a-6ab4-9643570916d2');
INSERT INTO public.events VALUES ('f21db000-161a-f7e8-4417-0c6ec018fc07', 'Title Notice 116', 'Description Notice 116', '2023-08-20 22:01:26+03', '2023-08-20 22:21:26+03', '2023-08-20 21:51:31+03', '42077d1c-ac74-a24e-69f5-f99423b5ee3e');
INSERT INTO public.events VALUES ('cfa22a34-b834-3480-1465-050b1d19d793', 'Title Notice 42', 'Description Notice 42', '2023-08-20 22:12:17+03', '2023-08-20 22:32:17+03', '2023-08-20 22:02:22+03', '3f24a0fb-3c3b-c0db-88bc-92361c2e7d83');
INSERT INTO public.events VALUES ('3a294436-6982-7e37-641e-cee2b3425725', 'Title Notice 44', 'Description Notice 44', '2023-08-20 22:12:19+03', '2023-08-20 22:32:19+03', '2023-08-20 22:02:24+03', 'bf871b64-7d55-7c80-20ae-e2b8088e6c73');
INSERT INTO public.events VALUES ('0d7dc7ee-ff59-ce26-28fc-f1e007734739', 'Title Notice 46', 'Description Notice 46', '2023-08-20 22:12:21+03', '2023-08-20 22:32:21+03', '2023-08-20 22:02:26+03', 'a77a5de5-649b-629c-9c41-2e7c3f9937b9');
INSERT INTO public.events VALUES ('540ae965-b404-e4b0-2682-d0a18aecec30', 'Title Notice 48', 'Description Notice 48', '2023-08-20 22:12:23+03', '2023-08-20 22:32:23+03', '2023-08-20 22:02:28+03', 'b7776eca-ad28-f694-a053-dc1db0244c1e');
INSERT INTO public.events VALUES ('28a315c9-baa1-621c-a732-2d8fefc77e55', 'Title Notice 50', 'Description Notice 50', '2023-08-20 22:12:25+03', '2023-08-20 22:32:25+03', '2023-08-20 22:02:30+03', 'da0b8809-7512-a69b-acae-e62c0c6efe54');
INSERT INTO public.events VALUES ('115bdced-89f3-2c77-8649-5e18440b9824', 'Title Notice 52', 'Description Notice 52', '2023-08-20 22:12:27+03', '2023-08-20 22:32:27+03', '2023-08-20 22:02:32+03', '2a633f28-fab6-266a-325d-904a92014c52');
INSERT INTO public.events VALUES ('a9fd708e-b69f-b68a-9011-b2af1369c7fd', 'Title Notice 54', 'Description Notice 54', '2023-08-20 22:12:29+03', '2023-08-20 22:32:29+03', '2023-08-20 22:02:34+03', 'ee2c3d1a-e88e-dd40-045b-87d24b8ab956');
INSERT INTO public.events VALUES ('f11d2d47-0767-4f98-d630-9f84d070cfa3', 'Title Notice 56', 'Description Notice 56', '2023-08-20 22:12:31+03', '2023-08-20 22:32:31+03', '2023-08-20 22:02:36+03', 'afedb613-f691-6e2e-854e-b0edb2eb7ad6');
INSERT INTO public.events VALUES ('0865d9d4-c9da-8174-4a67-7600ff866692', 'Title Notice 58', 'Description Notice 58', '2023-08-20 22:12:33+03', '2023-08-20 22:32:33+03', '2023-08-20 22:02:38+03', 'b1a52921-22ea-c7bf-baf8-a93e3772e671');
INSERT INTO public.events VALUES ('62d5497c-27a8-dfee-9dc8-0b6b73a2e962', 'Title Notice 60', 'Description Notice 60', '2023-08-20 22:12:35+03', '2023-08-20 22:32:35+03', '2023-08-20 22:02:40+03', 'e4003d2b-51da-0a57-d7d1-7ceb4e075f9c');
INSERT INTO public.events VALUES ('4947e1c7-942e-87c0-4673-ed4ada2f611f', 'Title Notice 62', 'Description Notice 62', '2023-08-20 22:12:37+03', '2023-08-20 22:32:37+03', '2023-08-20 22:02:42+03', 'fb0d84f7-2010-cde7-e2be-bbf48d2c14dd');
INSERT INTO public.events VALUES ('bbdf44e2-bc9b-afce-a340-313cb69d8c15', 'Title Notice 118', 'Description Notice 118', '2023-08-20 22:01:28+03', '2023-08-20 22:21:28+03', '2023-08-20 21:51:33+03', '96bbabdf-9bbb-4d20-2127-47d22b028a09');
INSERT INTO public.events VALUES ('afd05c23-96b9-e4fb-1143-db0a48c48574', 'Title Notice 120', 'Description Notice 120', '2023-08-20 22:01:30+03', '2023-08-20 22:21:30+03', '2023-08-20 21:51:35+03', '93959f7d-3d2a-cc10-eabc-eba9fa141564');
INSERT INTO public.events VALUES ('6315c883-5a08-7cbf-d7ad-3043f88860e6', 'Title Notice 64', 'Description Notice 64', '2023-08-20 22:12:39+03', '2023-08-20 22:32:39+03', '2023-08-20 22:02:44+03', 'bec9e7d0-22d2-55bf-b46f-604b19e7367d');
INSERT INTO public.events VALUES ('3860229f-1325-ad56-302c-200e98622d44', 'Title Notice 122', 'Description Notice 122', '2023-08-20 22:01:32+03', '2023-08-20 22:21:32+03', '2023-08-20 21:51:37+03', '335e98bd-eb2e-b026-df49-acccd6953568');
INSERT INTO public.events VALUES ('06c612f1-cc36-aad8-8ff6-5fc91d4c1f81', 'Title Notice 124', 'Description Notice 124', '2023-08-20 22:01:34+03', '2023-08-20 22:21:34+03', '2023-08-20 21:51:39+03', '9dd64969-84e7-ef2c-d2a8-eac5210cd0e6');
INSERT INTO public.events VALUES ('e00ec71d-d094-4cfe-c89b-4dced84f5a33', 'Title Notice 66', 'Description Notice 66', '2023-08-20 22:12:41+03', '2023-08-20 22:32:41+03', '2023-08-20 22:02:46+03', 'ef465f3c-d460-4696-5259-d75761539fe7');
INSERT INTO public.events VALUES ('7a1eba83-28fe-e817-a42a-8bdb5844bb6c', 'Title Notice 126', 'Description Notice 126', '2023-08-20 22:01:36+03', '2023-08-20 22:21:36+03', '2023-08-20 21:51:41+03', '17a95fa6-335a-4323-1b30-5e222b9e20cf');
INSERT INTO public.events VALUES ('d48001a4-3b25-f341-64c9-d98246872d94', 'Title Notice 128', 'Description Notice 128', '2023-08-20 22:01:38+03', '2023-08-20 22:21:38+03', '2023-08-20 21:51:43+03', '20dd7d75-f6c2-3f5a-3b8f-462ef5bb7fe8');
INSERT INTO public.events VALUES ('30c89499-179e-3365-8501-03463a349d11', 'Title Notice 68', 'Description Notice 68', '2023-08-20 22:12:43+03', '2023-08-20 22:32:43+03', '2023-08-20 22:02:48+03', 'baaa1aa5-ebde-af62-1ef5-141ab21dd7cd');
INSERT INTO public.events VALUES ('f2707b77-a30f-adbf-234d-91465d01601c', 'Title Notice 130', 'Description Notice 130', '2023-08-20 22:01:40+03', '2023-08-20 22:21:40+03', '2023-08-20 21:51:45+03', 'edd848c0-9dd3-bbb8-bfc5-057ae81c6e9a');
INSERT INTO public.events VALUES ('277ddd99-3fff-9b61-ce92-c119d9f5b2b9', 'Title Notice 132', 'Description Notice 132', '2023-08-20 22:01:42+03', '2023-08-20 22:21:42+03', '2023-08-20 21:51:47+03', '169706c2-2bb2-9c02-e403-e3ce0ca97136');
INSERT INTO public.events VALUES ('ad7d01f1-d258-ac12-ef37-df45a1c2fc2b', 'Title Notice 70', 'Description Notice 70', '2023-08-20 22:12:45+03', '2023-08-20 22:32:45+03', '2023-08-20 22:02:50+03', '97c52bc9-cc8a-a200-beff-46bf70f1a12d');
INSERT INTO public.events VALUES ('5f1134ff-7eea-cfaa-d49e-6d4520f77f1f', 'Title Notice 134', 'Description Notice 134', '2023-08-20 22:01:44+03', '2023-08-20 22:21:44+03', '2023-08-20 21:51:49+03', 'a48d2997-bb5a-db39-ba4a-33f51ef85750');
INSERT INTO public.events VALUES ('3f42fb0b-db2b-da27-221f-462f67e4c1b2', 'Title Notice 136', 'Description Notice 136', '2023-08-20 22:01:46+03', '2023-08-20 22:21:46+03', '2023-08-20 21:51:51+03', 'c062db7c-61a1-2438-3258-42f252e4c544');
INSERT INTO public.events VALUES ('e13e4202-8d3a-d676-3e4d-a59949f7d5c1', 'Title Notice 72', 'Description Notice 72', '2023-08-20 22:12:47+03', '2023-08-20 22:32:47+03', '2023-08-20 22:02:52+03', '3a442759-c4fa-b9ba-6213-0c27f785c999');
INSERT INTO public.events VALUES ('0c5898a7-b7c6-1888-1741-9e98b565b4d4', 'Title Notice 138', 'Description Notice 138', '2023-08-20 22:01:48+03', '2023-08-20 22:21:48+03', '2023-08-20 21:51:53+03', 'afee4840-2beb-7060-6535-0d70e4425fa6');
INSERT INTO public.events VALUES ('850b838b-1ade-363c-db60-8ece77853543', 'Title Notice 140', 'Description Notice 140', '2023-08-20 22:01:50+03', '2023-08-20 22:21:50+03', '2023-08-20 21:51:55+03', 'a871c6d1-db4e-c62e-0ce9-9d3e2500d93c');
INSERT INTO public.events VALUES ('7026bfff-3ddd-15b3-d5ec-adf1e20db38e', 'Title Notice 74', 'Description Notice 74', '2023-08-20 22:12:49+03', '2023-08-20 22:32:49+03', '2023-08-20 22:02:54+03', 'fff4ef40-3f11-40ff-570e-7f35cf70581b');
INSERT INTO public.events VALUES ('3caafdb8-0be4-f2d0-4ce1-77db57b4761e', 'Title Notice 142', 'Description Notice 142', '2023-08-20 22:01:52+03', '2023-08-20 22:21:52+03', '2023-08-20 21:51:57+03', 'aa590919-ab14-d091-ea6a-cf67e4f076e0');
INSERT INTO public.events VALUES ('971a2674-b9b3-d67d-d6af-e42c6a3864f5', 'Title Notice 144', 'Description Notice 144', '2023-08-20 22:01:54+03', '2023-08-20 22:21:54+03', '2023-08-20 21:51:59+03', '386331f6-4234-29b7-3d0a-1d52bd75ff3e');
INSERT INTO public.events VALUES ('9724bc11-713c-b2ca-1d21-8dde8ae3998a', 'Title Notice 76', 'Description Notice 76', '2023-08-20 22:12:51+03', '2023-08-20 22:32:51+03', '2023-08-20 22:02:56+03', 'eb1a696f-634b-7956-7547-cb70e04bf4e1');
INSERT INTO public.events VALUES ('9d576364-22ae-7808-98f7-cbc4046305ad', 'Title Notice 146', 'Description Notice 146', '2023-08-20 22:01:56+03', '2023-08-20 22:21:56+03', '2023-08-20 21:52:01+03', '5c798790-40a7-567d-0149-a76553bda891');
INSERT INTO public.events VALUES ('49b3dd80-ed45-1265-0983-fd5370575a95', 'Title Notice 148', 'Description Notice 148', '2023-08-20 22:01:58+03', '2023-08-20 22:21:58+03', '2023-08-20 21:52:03+03', 'cde5d250-b1ca-5283-537d-24933f4ef48b');
INSERT INTO public.events VALUES ('9ee85f87-d505-b9fd-0690-4b512000d941', 'Title Notice 78', 'Description Notice 78', '2023-08-20 22:12:53+03', '2023-08-20 22:32:53+03', '2023-08-20 22:02:58+03', '53694600-09b2-795a-58b5-592fef55ca14');
INSERT INTO public.events VALUES ('88340762-83d8-67c1-58c8-756df49308f1', 'Title Notice 150', 'Description Notice 150', '2023-08-20 22:02:00+03', '2023-08-20 22:22:00+03', '2023-08-20 21:52:05+03', 'cb00b5e4-d000-6a43-708c-c926590dc852');
INSERT INTO public.events VALUES ('6ecefba2-0200-dc5a-7053-9c9fcde48c1c', 'Title Notice 152', 'Description Notice 152', '2023-08-20 22:02:02+03', '2023-08-20 22:22:02+03', '2023-08-20 21:52:07+03', '46da030f-4086-f350-42ce-f230e564b3a2');
INSERT INTO public.events VALUES ('3ca1193e-cf93-8b3a-3512-e4173facb2b1', 'Title Notice 80', 'Description Notice 80', '2023-08-20 22:12:55+03', '2023-08-20 22:32:55+03', '2023-08-20 22:03:00+03', '6cbfd5ee-9771-dcfd-6cc6-aac575cb3348');
INSERT INTO public.events VALUES ('d46354be-de4e-636c-0c44-40f802589b96', 'Title Notice 154', 'Description Notice 154', '2023-08-20 22:02:04+03', '2023-08-20 22:22:04+03', '2023-08-20 21:52:09+03', '77b2116f-95f0-3c3a-9dd9-9f502bdf87f9');
INSERT INTO public.events VALUES ('4d6b5cf6-adf3-2fb5-75ea-ab485ccb3b91', 'Title Notice 156', 'Description Notice 156', '2023-08-20 22:02:06+03', '2023-08-20 22:22:06+03', '2023-08-20 21:52:11+03', '390e3ddf-1d07-178e-a2c6-727c46399d00');
INSERT INTO public.events VALUES ('a3e67061-13e2-4ac9-d406-8fce613d9fe3', 'Title Notice 82', 'Description Notice 82', '2023-08-20 22:12:57+03', '2023-08-20 22:32:57+03', '2023-08-20 22:03:02+03', '0f7d9309-a11e-a819-4275-1d6d1721d8a7');
INSERT INTO public.events VALUES ('1298645f-48d4-f144-347b-c6ef2c3538d0', 'Title Notice 158', 'Description Notice 158', '2023-08-20 22:02:08+03', '2023-08-20 22:22:08+03', '2023-08-20 21:52:13+03', '4dbfa7de-ece5-28bb-00fb-6f11823904fd');
INSERT INTO public.events VALUES ('df649184-5408-014a-5a15-fb4ed0d2c9c9', 'Title Notice 160', 'Description Notice 160', '2023-08-20 22:02:10+03', '2023-08-20 22:22:10+03', '2023-08-20 21:52:15+03', '2fe3f5e3-19e7-05c4-ddb4-4226a7b1f744');
INSERT INTO public.events VALUES ('c35f3810-58bb-913d-dd09-b0142db8eb24', 'Title Notice 84', 'Description Notice 84', '2023-08-20 22:12:59+03', '2023-08-20 22:32:59+03', '2023-08-20 22:03:04+03', 'af67e90b-6653-6b10-6d8e-ccdf56cc7a92');
INSERT INTO public.events VALUES ('6682154d-84cb-e39b-fde4-62e9bc9518bf', 'Title Notice 162', 'Description Notice 162', '2023-08-20 22:02:12+03', '2023-08-20 22:22:12+03', '2023-08-20 21:52:17+03', 'd333fd03-0c28-5898-bd67-d5ff05f33c74');
INSERT INTO public.events VALUES ('dd0f6733-d83e-bf4f-a173-eb4b8eec7931', 'Title Notice 164', 'Description Notice 164', '2023-08-20 22:02:14+03', '2023-08-20 22:22:14+03', '2023-08-20 21:52:19+03', 'd4934334-6347-8f6e-b0c0-e44fa1f136f1');
INSERT INTO public.events VALUES ('eb8454d7-230e-a352-c1c9-99d19f917d7a', 'Title Notice 86', 'Description Notice 86', '2023-08-20 22:13:01+03', '2023-08-20 22:33:01+03', '2023-08-20 22:03:06+03', 'bcdcad66-a489-fe19-539e-867991c44c88');
INSERT INTO public.events VALUES ('ed2ac41c-3165-26fa-956b-8dda5719dbfe', 'Title Notice 88', 'Description Notice 88', '2023-08-20 22:13:03+03', '2023-08-20 22:33:03+03', '2023-08-20 22:03:08+03', '190b84f3-8a36-0f5c-040f-59c999626172');
INSERT INTO public.events VALUES ('c950144b-48c1-3bb8-cea6-5e23e28e3d5f', 'Title Notice 90', 'Description Notice 90', '2023-08-20 22:13:05+03', '2023-08-20 22:33:05+03', '2023-08-20 22:03:10+03', '8a496b26-052f-1793-5eb7-834b3027e167');
INSERT INTO public.events VALUES ('95949310-9cf6-3a61-2ad4-817a50a42109', 'Title Notice 92', 'Description Notice 92', '2023-08-20 22:13:07+03', '2023-08-20 22:33:07+03', '2023-08-20 22:03:12+03', '9304ca6f-dc3b-6454-68f6-f85a7f33490f');
INSERT INTO public.events VALUES ('9ca0f4ea-ddd0-e05f-3175-cf24f18c9a78', 'Title Notice 94', 'Description Notice 94', '2023-08-20 22:13:09+03', '2023-08-20 22:33:09+03', '2023-08-20 22:03:14+03', '217b8d12-4327-ba35-2073-30e4c4945825');
INSERT INTO public.events VALUES ('798e316b-cfef-56ce-96bb-1fe8db1a2755', 'Title Notice 96', 'Description Notice 96', '2023-08-20 22:13:11+03', '2023-08-20 22:33:11+03', '2023-08-20 22:03:16+03', '2cca4763-f798-17a0-2d0a-a98fad204ca5');
INSERT INTO public.events VALUES ('5f07917e-e797-cee3-3842-0926e09d5018', 'Title Notice 98', 'Description Notice 98', '2023-08-20 22:13:13+03', '2023-08-20 22:33:13+03', '2023-08-20 22:03:18+03', '54a04abe-5fd3-c67c-2e27-30b0a3eb0155');
INSERT INTO public.events VALUES ('fc0c8e81-a8cc-382c-19a7-123179a1425a', 'Title Notice 100', 'Description Notice 100', '2023-08-20 22:13:15+03', '2023-08-20 22:33:15+03', '2023-08-20 22:03:20+03', '977a5378-acdb-b749-013a-06cf3c79de68');
INSERT INTO public.events VALUES ('2921e7f7-38e9-fe96-b17f-aba700ed6bac', 'Title Notice 102', 'Description Notice 102', '2023-08-20 22:13:17+03', '2023-08-20 22:33:17+03', '2023-08-20 22:03:22+03', '81a6b0e7-eee1-b36f-2780-675faf3f751d');
INSERT INTO public.events VALUES ('5cca0472-8478-fa40-5e8c-9a5006272dbf', 'Title Notice 104', 'Description Notice 104', '2023-08-20 22:13:19+03', '2023-08-20 22:33:19+03', '2023-08-20 22:03:24+03', '1e75b375-20d6-96b0-cd91-81bb163cc8cb');
INSERT INTO public.events VALUES ('39c00cbf-15b8-f62c-a1dc-d35a5bfed21f', 'Title Notice 2', 'Description Notice 2', '2023-08-21 11:04:24+03', '2023-08-21 11:24:24+03', '2023-08-21 10:54:29+03', '56231a34-d8c4-fea6-744b-5514f51f71ec');
INSERT INTO public.events VALUES ('a12ec0a7-9620-443f-e0d2-218fbd4cf520', 'Title Notice 166', 'Description Notice 166', '2023-08-20 22:02:16+03', '2023-08-20 22:22:16+03', '2023-08-20 21:52:21+03', 'f9c7336e-23cf-b629-c4a7-e35b005f2151');
INSERT INTO public.events VALUES ('896950ed-b6cf-b213-aabf-f8abda7ff01c', 'Title Notice 106', 'Description Notice 106', '2023-08-20 22:13:21+03', '2023-08-20 22:33:21+03', '2023-08-20 22:03:26+03', 'a5178bfb-5ef5-4b6b-8d31-d7fee7925444');
INSERT INTO public.events VALUES ('73f865a0-17b0-2b8b-74f5-2c55bca60c4d', 'Title Notice 168', 'Description Notice 168', '2023-08-20 22:02:18+03', '2023-08-20 22:22:18+03', '2023-08-20 21:52:23+03', 'a07b81dc-ead0-e539-8f40-93a2417cf763');
INSERT INTO public.events VALUES ('2e52cd50-3240-c4ab-355b-efeff71a832e', 'Title Notice 170', 'Description Notice 170', '2023-08-20 22:02:20+03', '2023-08-20 22:22:20+03', '2023-08-20 21:52:25+03', 'fa09e409-4ed7-12dd-a0da-c98e2f08d65e');
INSERT INTO public.events VALUES ('5f4926b2-f4e0-a2d2-c9bf-51f06193e0aa', 'Title Notice 108', 'Description Notice 108', '2023-08-20 22:13:23+03', '2023-08-20 22:33:23+03', '2023-08-20 22:03:28+03', '07f17e9a-81d9-9239-a8aa-51ec24609454');
INSERT INTO public.events VALUES ('6b6c9325-fb84-89e9-ddec-4edd2590fe29', 'Title Notice 172', 'Description Notice 172', '2023-08-20 22:02:22+03', '2023-08-20 22:22:22+03', '2023-08-20 21:52:27+03', 'ac92ab29-8413-9767-231b-ba95e09c95a9');
INSERT INTO public.events VALUES ('afdc30b8-f0bd-a689-dc69-055123461046', 'Title Notice 174', 'Description Notice 174', '2023-08-20 22:02:24+03', '2023-08-20 22:22:24+03', '2023-08-20 21:52:29+03', '7efc5968-1406-1842-4bff-9c5936857b4c');
INSERT INTO public.events VALUES ('00bbdcf8-1213-52e9-91ea-bad30e873ef5', 'Title Notice 110', 'Description Notice 110', '2023-08-20 22:13:25+03', '2023-08-20 22:33:25+03', '2023-08-20 22:03:30+03', '6b843aa9-5bff-2ec5-c7d5-a61e70f28a5c');
INSERT INTO public.events VALUES ('e686f70c-d7c4-b26b-0f3d-762245e65c28', 'Title Notice 176', 'Description Notice 176', '2023-08-20 22:02:26+03', '2023-08-20 22:22:26+03', '2023-08-20 21:52:31+03', 'eb1a2c62-9f85-43ce-190f-f893e4d38788');
INSERT INTO public.events VALUES ('83b02d4a-e362-e10a-cfbe-7b829591cfb1', 'Title Notice 178', 'Description Notice 178', '2023-08-20 22:02:28+03', '2023-08-20 22:22:28+03', '2023-08-20 21:52:33+03', 'd0dbb55b-8b43-b2a0-7cb4-276be04771bd');
INSERT INTO public.events VALUES ('4d642ad8-065c-2853-121f-1e4cd6ff0f2e', 'Title Notice 112', 'Description Notice 112', '2023-08-20 22:13:27+03', '2023-08-20 22:33:27+03', '2023-08-20 22:03:32+03', '28ba1736-e5dc-223e-1dac-3f524eab90fe');
INSERT INTO public.events VALUES ('7d26097e-14b4-7fb1-5bf2-a980b928f023', 'Title Notice 180', 'Description Notice 180', '2023-08-20 22:02:30+03', '2023-08-20 22:22:30+03', '2023-08-20 21:52:35+03', '76c39960-c17e-0fc0-ab64-fdaf699baa94');
INSERT INTO public.events VALUES ('a485cb2e-7e56-62c8-1ff3-ccaadc635e0a', 'Title Notice 182', 'Description Notice 182', '2023-08-20 22:02:32+03', '2023-08-20 22:22:32+03', '2023-08-20 21:52:37+03', 'b19ba3eb-d6e5-2d9e-a111-0e50b48649a8');
INSERT INTO public.events VALUES ('0573ace6-0e06-0b04-01a6-428d1b1ea34b', 'Title Notice 114', 'Description Notice 114', '2023-08-20 22:13:29+03', '2023-08-20 22:33:29+03', '2023-08-20 22:03:34+03', '2d27709c-7ea1-30f0-d26f-b3e93c0ebee1');
INSERT INTO public.events VALUES ('68ddaef3-8739-7790-baea-1ef4ce931dd7', 'Title Notice 184', 'Description Notice 184', '2023-08-20 22:02:34+03', '2023-08-20 22:22:34+03', '2023-08-20 21:52:39+03', '02bb5f6f-be9d-b1d3-f63c-e7630af2921c');
INSERT INTO public.events VALUES ('591941a8-3902-12f7-85cf-d4ebd214e70b', 'Title Notice 186', 'Description Notice 186', '2023-08-20 22:02:36+03', '2023-08-20 22:22:36+03', '2023-08-20 21:52:41+03', 'd291dd82-ac88-03fe-9277-e5914fb59da3');
INSERT INTO public.events VALUES ('303ae14a-0142-41ff-f698-eca10a0e15b6', 'Title Notice 116', 'Description Notice 116', '2023-08-20 22:13:31+03', '2023-08-20 22:33:31+03', '2023-08-20 22:03:36+03', 'fcf2a32e-6397-5ec1-3fba-a073bd24fa00');
INSERT INTO public.events VALUES ('41a1b2d3-8e13-db54-8e29-047033724c47', 'Title Notice 188', 'Description Notice 188', '2023-08-20 22:02:38+03', '2023-08-20 22:22:38+03', '2023-08-20 21:52:43+03', '736474c5-92df-9487-be7b-a6cac1b6dc2d');
INSERT INTO public.events VALUES ('7fdc1eef-7e11-acc3-0498-97ca2e219895', 'Title Notice 190', 'Description Notice 190', '2023-08-20 22:02:40+03', '2023-08-20 22:22:40+03', '2023-08-20 21:52:45+03', '4141d3cd-9177-ff06-d787-db61e683267f');
INSERT INTO public.events VALUES ('ea0a4b5d-59b8-7220-c48c-bf1684f47243', 'Title Notice 118', 'Description Notice 118', '2023-08-20 22:13:33+03', '2023-08-20 22:33:33+03', '2023-08-20 22:03:38+03', '0cbcddf3-3eb2-e2b6-668e-5607f56ab45f');
INSERT INTO public.events VALUES ('4a34e687-04b5-0a1a-8996-7572fc14b39d', 'Title Notice 192', 'Description Notice 192', '2023-08-20 22:02:42+03', '2023-08-20 22:22:42+03', '2023-08-20 21:52:47+03', '98751b2b-fd5e-b83a-f8fa-c099a22967b8');
INSERT INTO public.events VALUES ('f1e7fb84-6d86-35f9-3a43-d8a59988bd75', 'Title Notice 194', 'Description Notice 194', '2023-08-20 22:02:44+03', '2023-08-20 22:22:44+03', '2023-08-20 21:52:49+03', '828ed0d8-db0f-429b-b722-08e6e770fe3d');
INSERT INTO public.events VALUES ('31ac7a03-bd0b-bd14-dacc-e6631c090216', 'Title Notice 120', 'Description Notice 120', '2023-08-20 22:13:35+03', '2023-08-20 22:33:35+03', '2023-08-20 22:03:40+03', 'bfdeae88-3267-7e7a-1bf5-76837396c76c');
INSERT INTO public.events VALUES ('9d90798f-ba5e-455d-df25-ad428b07b91d', 'Title Notice 196', 'Description Notice 196', '2023-08-20 22:02:46+03', '2023-08-20 22:22:46+03', '2023-08-20 21:52:51+03', '72974e84-41f9-e806-7846-f23e9eb13050');
INSERT INTO public.events VALUES ('85ee4aed-aee8-ea9e-ed81-a2c32372c994', 'Title Notice 198', 'Description Notice 198', '2023-08-20 22:02:48+03', '2023-08-20 22:22:48+03', '2023-08-20 21:52:53+03', '9a564f91-28eb-9247-724e-0be7771e749b');
INSERT INTO public.events VALUES ('0ff58969-9b6d-4d61-54aa-b6fc5543261c', 'Title Notice 4', 'Description Notice 4', '2023-08-21 11:04:26+03', '2023-08-21 11:24:26+03', '2023-08-21 10:54:31+03', '3a9f0b0c-476c-2cf6-4c1f-4fe92452017f');
INSERT INTO public.events VALUES ('58be0020-e839-a65b-9ca7-7735186dd1d9', 'Title Notice 6', 'Description Notice 6', '2023-08-21 11:04:28+03', '2023-08-21 11:24:28+03', '2023-08-21 10:54:33+03', '1a35d55a-56b8-067b-0b21-452993e46d0f');
INSERT INTO public.events VALUES ('b7788887-29bc-53cf-c992-f1c0dc68b38a', 'Title Notice 8', 'Description Notice 8', '2023-08-21 11:04:30+03', '2023-08-21 11:24:30+03', '2023-08-21 10:54:35+03', 'e416f570-963b-eb62-8785-63954197cd4e');
INSERT INTO public.events VALUES ('069391a5-9a0f-77ff-fa3e-7bd83615fd42', 'Title Notice 10', 'Description Notice 10', '2023-08-21 11:04:32+03', '2023-08-21 11:24:32+03', '2023-08-21 10:54:37+03', '6ba45436-6a15-9ac4-7076-4cf81fdd63b9');
INSERT INTO public.events VALUES ('0cd04c85-a0ab-c3bd-4adb-8d0976600088', 'Title Notice 12', 'Description Notice 12', '2023-08-21 11:04:34+03', '2023-08-21 11:24:34+03', '2023-08-21 10:54:39+03', '19397e74-8964-f17e-b711-1285de1073d6');
INSERT INTO public.events VALUES ('ed490ece-5244-2e16-a544-9390a0cedbfa', 'Title Notice 14', 'Description Notice 14', '2023-08-21 11:04:36+03', '2023-08-21 11:24:36+03', '2023-08-21 10:54:41+03', 'a11d5195-5f45-a004-dca2-7956a86950a6');
INSERT INTO public.events VALUES ('369c28ee-894a-fbff-6d12-e5b99fa64bee', 'Title Notice 16', 'Description Notice 16', '2023-08-21 11:04:38+03', '2023-08-21 11:24:38+03', '2023-08-21 10:54:43+03', '3f1ac867-cc90-4783-a80f-dd180316452c');
INSERT INTO public.events VALUES ('30ae887b-7249-7b38-6925-0b5512a0cfe4', 'Title Notice 18', 'Description Notice 18', '2023-08-21 11:04:40+03', '2023-08-21 11:24:40+03', '2023-08-21 10:54:45+03', 'e0d05ae6-5ab4-8bda-250f-ef90fb4cb1c2');
INSERT INTO public.events VALUES ('e038ba95-f297-dcf2-db9e-7ba335516de8', 'Title Notice 20', 'Description Notice 20', '2023-08-21 11:04:42+03', '2023-08-21 11:24:42+03', '2023-08-21 10:54:47+03', '4337fa2a-ed54-797a-605f-d3b04e187f17');
INSERT INTO public.events VALUES ('5a441425-3295-b337-96e3-0915658a6efe', 'Title Notice 22', 'Description Notice 22', '2023-08-21 11:04:44+03', '2023-08-21 11:24:44+03', '2023-08-21 10:54:49+03', '219f5519-d11f-1fdd-5f7d-9537c0a20301');
INSERT INTO public.events VALUES ('668960c9-9718-4bc7-a4bb-f6650aae5366', 'Title Notice 24', 'Description Notice 24', '2023-08-21 11:04:46+03', '2023-08-21 11:24:46+03', '2023-08-21 10:54:51+03', '3a5b9275-51d6-8f51-237d-fdbc98ef385e');
INSERT INTO public.events VALUES ('c4a79fdb-1001-1447-cc8a-f4b2a1daa3e4', 'Title Notice 26', 'Description Notice 26', '2023-08-21 11:04:48+03', '2023-08-21 11:24:48+03', '2023-08-21 10:54:53+03', '7ee35bcc-cca0-fad4-216a-2eb4c894e202');
INSERT INTO public.events VALUES ('57f834f8-e326-da71-5891-b0c081e25cc3', 'Title Notice 28', 'Description Notice 28', '2023-08-21 11:04:50+03', '2023-08-21 11:24:50+03', '2023-08-21 10:54:55+03', '887784ff-4682-6732-ea85-fffcdee4d6b0');
INSERT INTO public.events VALUES ('69c1f952-b728-6c0c-a9d5-f0c3973fa0c8', 'Title Notice 30', 'Description Notice 30', '2023-08-21 11:04:52+03', '2023-08-21 11:24:52+03', '2023-08-21 10:54:57+03', 'cddee941-4ade-880b-b79f-4d418ededca1');
INSERT INTO public.events VALUES ('ff5d4e32-59dc-8073-c222-4e53ac006e2d', 'Title Notice 32', 'Description Notice 32', '2023-08-21 11:04:54+03', '2023-08-21 11:24:54+03', '2023-08-21 10:54:59+03', '40a99d89-2aec-37c0-cd99-4e637135c447');
INSERT INTO public.events VALUES ('570c2dd1-3e20-a615-384e-6e5fd9847703', 'Title Notice 34', 'Description Notice 34', '2023-08-21 11:04:56+03', '2023-08-21 11:24:56+03', '2023-08-21 10:55:01+03', 'f2e46bf7-e552-298e-b326-3c09bce8bd16');
INSERT INTO public.events VALUES ('626a3e6d-32ea-25ba-9366-bc15ddded3ef', 'Title Notice 36', 'Description Notice 36', '2023-08-21 11:04:58+03', '2023-08-21 11:24:58+03', '2023-08-21 10:55:03+03', '163707a5-2835-c6f1-3ddc-c280435320f4');
INSERT INTO public.events VALUES ('bab1fa62-59ad-3041-bc7a-958f49bfa982', 'Title Notice 38', 'Description Notice 38', '2023-08-21 11:05:00+03', '2023-08-21 11:25:00+03', '2023-08-21 10:55:05+03', '28bfe143-e885-8cfe-17b8-bf5c44165d5f');
INSERT INTO public.events VALUES ('8e22947b-19aa-067e-120d-c27b3209f393', 'Title Notice 40', 'Description Notice 40', '2023-08-21 11:05:02+03', '2023-08-21 11:25:02+03', '2023-08-21 10:55:07+03', '86743643-0cd2-2711-42a6-ee354cfefcc4');
INSERT INTO public.events VALUES ('4c1a8ccc-a653-c19d-ba1d-c4eb11a5c00f', 'Title Notice 42', 'Description Notice 42', '2023-08-21 11:05:04+03', '2023-08-21 11:25:04+03', '2023-08-21 10:55:09+03', '9f0921f4-796c-732f-ecff-b9bae016ce8b');
INSERT INTO public.events VALUES ('8ae0c937-9ed7-95c9-3538-f96ada907b6e', 'Title Notice 44', 'Description Notice 44', '2023-08-21 11:05:06+03', '2023-08-21 11:25:06+03', '2023-08-21 10:55:11+03', '4b1b17b2-f651-6197-52b3-cb537fe75f53');
INSERT INTO public.events VALUES ('ae93ab14-fe99-18cf-e74e-4e1b1db3b825', 'Title Notice 46', 'Description Notice 46', '2023-08-21 11:05:08+03', '2023-08-21 11:25:08+03', '2023-08-21 10:55:13+03', '121bce0b-8f2d-945b-f868-e1041c6f51fa');
INSERT INTO public.events VALUES ('033dbb90-e698-e461-af48-79d44fc83747', 'Title Notice 48', 'Description Notice 48', '2023-08-21 11:05:10+03', '2023-08-21 11:25:10+03', '2023-08-21 10:55:15+03', '5153c7a3-cf63-1e70-a078-3b3bcfb23a9e');
INSERT INTO public.events VALUES ('34857bf0-51b1-0d8b-ebed-1b7020020d58', 'Title Notice 50', 'Description Notice 50', '2023-08-21 11:05:12+03', '2023-08-21 11:25:12+03', '2023-08-21 10:55:17+03', '3a32a530-9ba3-7378-819f-06940166012a');
INSERT INTO public.events VALUES ('cbaf95fb-661b-b541-c6af-99b6932a0293', 'Title Notice 52', 'Description Notice 52', '2023-08-21 11:05:14+03', '2023-08-21 11:25:14+03', '2023-08-21 10:55:19+03', 'c18d0a2d-07fd-6535-596c-07fa4cc2fda8');
INSERT INTO public.events VALUES ('d484659c-df11-2678-01c6-54614876ce03', 'Title Notice 54', 'Description Notice 54', '2023-08-21 11:05:16+03', '2023-08-21 11:25:16+03', '2023-08-21 10:55:21+03', '3fd00ec5-6d24-da67-84f2-658e58e75c5b');
INSERT INTO public.events VALUES ('7a9a1e5b-9e44-cace-7492-49ff01f3de8d', 'Title Notice 56', 'Description Notice 56', '2023-08-21 11:05:18+03', '2023-08-21 11:25:18+03', '2023-08-21 10:55:23+03', '71d5f9f7-b070-6a69-213c-1b90fbd97f8a');
INSERT INTO public.events VALUES ('9420f456-7171-a76b-90dc-c4619fd6c3b8', 'Title Notice 58', 'Description Notice 58', '2023-08-21 11:05:20+03', '2023-08-21 11:25:20+03', '2023-08-21 10:55:25+03', '8a228d92-2e7c-12d5-b14e-eb4bfa86acad');
INSERT INTO public.events VALUES ('dedd3688-1417-a90f-4875-9e06e6f781d4', 'Title Notice 60', 'Description Notice 60', '2023-08-21 11:05:22+03', '2023-08-21 11:25:22+03', '2023-08-21 10:55:27+03', '30a9984e-ebb6-7f40-6205-4bb736e37423');
INSERT INTO public.events VALUES ('c4c344fa-4c91-48d8-818e-cf6effa53212', 'Title Notice 62', 'Description Notice 62', '2023-08-21 11:05:24+03', '2023-08-21 11:25:24+03', '2023-08-21 10:55:29+03', '94457d81-9cd4-955c-4b02-754f2e217f2b');
INSERT INTO public.events VALUES ('3ba14cac-b318-448b-1478-a009f08c9bae', 'Title Notice 64', 'Description Notice 64', '2023-08-21 11:05:26+03', '2023-08-21 11:25:26+03', '2023-08-21 10:55:31+03', '8f3c4861-9f72-54c0-4c28-f318f9ae4ce1');
INSERT INTO public.events VALUES ('2e2b0001-dd92-d95b-02b2-b89461709695', 'Title Notice 66', 'Description Notice 66', '2023-08-21 11:05:28+03', '2023-08-21 11:25:28+03', '2023-08-21 10:55:33+03', '9b6f9daa-4756-e9a2-6f01-1437d19910bd');
INSERT INTO public.events VALUES ('665b22c5-656f-dc94-b3cf-f42b91353d62', 'Title Notice 68', 'Description Notice 68', '2023-08-21 11:05:30+03', '2023-08-21 11:25:30+03', '2023-08-21 10:55:35+03', '32c09fa5-8b31-c489-4695-d9c837e07ecb');
INSERT INTO public.events VALUES ('78bc3293-6ae6-2e4e-58c3-5b2c05d51720', 'Title Notice 70', 'Description Notice 70', '2023-08-21 11:05:32+03', '2023-08-21 11:25:32+03', '2023-08-21 10:55:37+03', 'b198080c-3d2d-e9cb-05e6-945733a41996');
INSERT INTO public.events VALUES ('860bd348-509e-cf5c-b4fb-a95ac92317cb', 'Title Notice 72', 'Description Notice 72', '2023-08-21 11:05:34+03', '2023-08-21 11:25:34+03', '2023-08-21 10:55:39+03', 'c3c18d66-caab-66d5-af69-2ca52a0dd06c');
INSERT INTO public.events VALUES ('83e94be2-1479-6768-c4df-4db6080642da', 'Title Notice 74', 'Description Notice 74', '2023-08-21 11:05:36+03', '2023-08-21 11:25:36+03', '2023-08-21 10:55:41+03', 'e2442a90-5e24-340c-f407-40aef7843891');
INSERT INTO public.events VALUES ('cc59e3f4-84d0-897a-7290-f423f98b744d', 'Title Notice 76', 'Description Notice 76', '2023-08-21 11:05:38+03', '2023-08-21 11:25:38+03', '2023-08-21 10:55:43+03', '1c253bcc-6cd2-b06e-eafd-c407c8daf26a');
INSERT INTO public.events VALUES ('c7874de0-2db9-4efa-cfbb-67b3246578ab', 'Title Notice 78', 'Description Notice 78', '2023-08-21 11:05:40+03', '2023-08-21 11:25:40+03', '2023-08-21 10:55:45+03', '0fecadbe-c11f-4947-d8ab-ad4a9035fe27');
INSERT INTO public.events VALUES ('32e60af0-2b24-7e11-02fd-49beeddeeea4', 'Title Notice 80', 'Description Notice 80', '2023-08-21 11:05:42+03', '2023-08-21 11:25:42+03', '2023-08-21 10:55:47+03', '4088af20-b219-88bb-c914-3532e99a3b46');
INSERT INTO public.events VALUES ('1b8afe58-a992-4961-b22f-fa2a093cddc8', 'Title Delete 81', 'Description Delete 81', '2022-08-21 10:55:44+03', '2022-08-21 11:25:44+03', '2022-08-21 11:55:44+03', '4de44de7-6f40-3bf3-55ed-c076bd61eb86');
INSERT INTO public.events VALUES ('e30e9825-66bd-085f-f0ef-6a2a80c34b1c', 'Title Notice 82', 'Description Notice 82', '2023-08-21 11:05:44+03', '2023-08-21 11:25:44+03', '2023-08-21 10:55:49+03', 'd424ac81-d822-e844-581e-8e01e9269fa3');
INSERT INTO public.events VALUES ('457533f4-a232-ae15-1848-9633ceb1350f', 'Title Delete 83', 'Description Delete 83', '2022-08-21 10:55:46+03', '2022-08-21 11:25:46+03', '2022-08-21 11:55:46+03', '0fb2a00f-7c5a-5d9c-fee7-31dda53a6ed9');
INSERT INTO public.events VALUES ('4cea97db-66c3-5f2b-a89c-a8b946ca907a', 'Title Notice 84', 'Description Notice 84', '2023-08-21 11:05:46+03', '2023-08-21 11:25:46+03', '2023-08-21 10:55:51+03', '151fab08-bcc0-29ba-f085-36e3b4a33276');


--
-- Data for Name: notes_check; Type: TABLE DATA; Schema: public; Owner: calendar
--

INSERT INTO public.notes_check VALUES ('2023-08-21 10:55:43.338984+03');


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: calendar
--

INSERT INTO public.schema_migrations VALUES (1, false);


--
-- Name: events events_pkey; Type: CONSTRAINT; Schema: public; Owner: calendar
--

ALTER TABLE ONLY public.events
    ADD CONSTRAINT events_pkey PRIMARY KEY (event_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: calendar
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: events_date_time_notice_date_time_start_index; Type: INDEX; Schema: public; Owner: calendar
--

CREATE INDEX events_date_time_notice_date_time_start_index ON public.events USING btree (date_time_notice, date_time_start);


--
-- Name: events_date_time_start_index; Type: INDEX; Schema: public; Owner: calendar
--

CREATE INDEX events_date_time_start_index ON public.events USING btree (date_time_start);


--
-- PostgreSQL database dump complete
--

