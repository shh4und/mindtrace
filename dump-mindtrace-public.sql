--
-- PostgreSQL database dump
--

-- Dumped from database version 17.6
-- Dumped by pg_dump version 17.5

-- Started on 2025-11-23 18:47:07 -03

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- TOC entry 237 (class 1259 OID 16552)
-- Name: atribuicoes; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.atribuicoes (
    id bigint NOT NULL,
    paciente_id bigint NOT NULL,
    instrumento_id bigint NOT NULL,
    status text DEFAULT 'PENDENTE'::text,
    data_atribuicao timestamp with time zone,
    data_resposta timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.atribuicoes OWNER TO admin;

--
-- TOC entry 236 (class 1259 OID 16551)
-- Name: atribuicoes_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.atribuicoes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.atribuicoes_id_seq OWNER TO admin;

--
-- TOC entry 3564 (class 0 OID 0)
-- Dependencies: 236
-- Name: atribuicoes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.atribuicoes_id_seq OWNED BY public.atribuicoes.id;


--
-- TOC entry 229 (class 1259 OID 16483)
-- Name: convites; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.convites (
    id bigint NOT NULL,
    profissional_id bigint NOT NULL,
    token text NOT NULL,
    data_expiracao timestamp with time zone NOT NULL,
    usado boolean DEFAULT false,
    paciente_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.convites OWNER TO admin;

--
-- TOC entry 228 (class 1259 OID 16482)
-- Name: convites_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.convites_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.convites_id_seq OWNER TO admin;

--
-- TOC entry 3565 (class 0 OID 0)
-- Dependencies: 228
-- Name: convites_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.convites_id_seq OWNED BY public.convites.id;


--
-- TOC entry 231 (class 1259 OID 16506)
-- Name: instrumentos; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.instrumentos (
    id bigint NOT NULL,
    codigo text NOT NULL,
    nome text NOT NULL,
    descricao text,
    algoritmo_pontuacao text NOT NULL,
    versao bigint DEFAULT 1,
    esta_ativo boolean DEFAULT true,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.instrumentos OWNER TO admin;

--
-- TOC entry 230 (class 1259 OID 16505)
-- Name: instrumentos_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.instrumentos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.instrumentos_id_seq OWNER TO admin;

--
-- TOC entry 3566 (class 0 OID 0)
-- Dependencies: 230
-- Name: instrumentos_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.instrumentos_id_seq OWNED BY public.instrumentos.id;


--
-- TOC entry 227 (class 1259 OID 16467)
-- Name: notificacoes; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.notificacoes (
    id bigint NOT NULL,
    usuario_id bigint NOT NULL,
    alerta_id bigint,
    conteudo text NOT NULL,
    status character varying(50) DEFAULT 'NAOLIDA'::character varying NOT NULL,
    data_envio timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


ALTER TABLE public.notificacoes OWNER TO admin;

--
-- TOC entry 226 (class 1259 OID 16466)
-- Name: notificacoes_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.notificacoes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.notificacoes_id_seq OWNER TO admin;

--
-- TOC entry 3567 (class 0 OID 0)
-- Dependencies: 226
-- Name: notificacoes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.notificacoes_id_seq OWNED BY public.notificacoes.id;


--
-- TOC entry 235 (class 1259 OID 16536)
-- Name: opcoes_escala; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.opcoes_escala (
    id bigint NOT NULL,
    instrumento_id bigint NOT NULL,
    valor bigint NOT NULL,
    rotulo text NOT NULL
);


ALTER TABLE public.opcoes_escala OWNER TO admin;

--
-- TOC entry 234 (class 1259 OID 16535)
-- Name: opcoes_escala_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.opcoes_escala_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.opcoes_escala_id_seq OWNER TO admin;

--
-- TOC entry 3568 (class 0 OID 0)
-- Dependencies: 234
-- Name: opcoes_escala_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.opcoes_escala_id_seq OWNED BY public.opcoes_escala.id;


--
-- TOC entry 222 (class 1259 OID 16418)
-- Name: pacientes; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.pacientes (
    id bigint NOT NULL,
    usuario_id bigint NOT NULL,
    data_nascimento timestamp with time zone,
    dependente boolean,
    nome_responsavel character varying(255),
    contato_responsavel character varying(11),
    data_inicio_tratamento timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.pacientes OWNER TO admin;

--
-- TOC entry 221 (class 1259 OID 16417)
-- Name: pacientes_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.pacientes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.pacientes_id_seq OWNER TO admin;

--
-- TOC entry 3569 (class 0 OID 0)
-- Dependencies: 221
-- Name: pacientes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.pacientes_id_seq OWNED BY public.pacientes.id;


--
-- TOC entry 233 (class 1259 OID 16519)
-- Name: perguntas; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.perguntas (
    id bigint NOT NULL,
    instrumento_id bigint NOT NULL,
    ordem_item bigint,
    dominio character varying(100),
    conteudo text NOT NULL,
    eh_pontuacao_invertida boolean DEFAULT false
);


ALTER TABLE public.perguntas OWNER TO admin;

--
-- TOC entry 232 (class 1259 OID 16518)
-- Name: perguntas_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.perguntas_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.perguntas_id_seq OWNER TO admin;

--
-- TOC entry 3570 (class 0 OID 0)
-- Dependencies: 232
-- Name: perguntas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.perguntas_id_seq OWNED BY public.perguntas.id;


--
-- TOC entry 220 (class 1259 OID 16401)
-- Name: profissionais; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.profissionais (
    id bigint NOT NULL,
    usuario_id bigint NOT NULL,
    data_nascimento timestamp with time zone,
    especialidade character varying(255) NOT NULL,
    registro_profissional character varying(12) NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.profissionais OWNER TO admin;

--
-- TOC entry 219 (class 1259 OID 16400)
-- Name: profissionais_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.profissionais_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.profissionais_id_seq OWNER TO admin;

--
-- TOC entry 3571 (class 0 OID 0)
-- Dependencies: 219
-- Name: profissionais_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.profissionais_id_seq OWNED BY public.profissionais.id;


--
-- TOC entry 223 (class 1259 OID 16432)
-- Name: profissional_paciente; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.profissional_paciente (
    paciente_id bigint NOT NULL,
    profissional_id bigint NOT NULL
);


ALTER TABLE public.profissional_paciente OWNER TO admin;

--
-- TOC entry 225 (class 1259 OID 16448)
-- Name: registros_humor; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.registros_humor (
    id bigint NOT NULL,
    paciente_id bigint NOT NULL,
    nivel_humor smallint NOT NULL,
    horas_sono smallint NOT NULL,
    nivel_energia smallint NOT NULL,
    nivel_stress smallint NOT NULL,
    auto_cuidado text NOT NULL,
    observacoes text,
    data_hora_registro timestamp with time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    created_at timestamp with time zone,
    CONSTRAINT chk_registros_humor_horas_sono CHECK (((horas_sono >= 0) AND (horas_sono <= 12))),
    CONSTRAINT chk_registros_humor_nivel_energia CHECK (((nivel_energia >= 1) AND (nivel_energia <= 10))),
    CONSTRAINT chk_registros_humor_nivel_humor CHECK (((nivel_humor >= 1) AND (nivel_humor <= 5))),
    CONSTRAINT chk_registros_humor_nivel_stress CHECK (((nivel_stress >= 1) AND (nivel_stress <= 10)))
);


ALTER TABLE public.registros_humor OWNER TO admin;

--
-- TOC entry 224 (class 1259 OID 16447)
-- Name: registros_humor_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.registros_humor_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.registros_humor_id_seq OWNER TO admin;

--
-- TOC entry 3572 (class 0 OID 0)
-- Dependencies: 224
-- Name: registros_humor_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.registros_humor_id_seq OWNED BY public.registros_humor.id;


--
-- TOC entry 239 (class 1259 OID 16576)
-- Name: respostas; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.respostas (
    id bigint NOT NULL,
    atribuicao_id bigint NOT NULL,
    pontuacao_total numeric(10,2),
    classificacao character varying(255),
    dados_brutos jsonb,
    data_resposta timestamp with time zone,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.respostas OWNER TO admin;

--
-- TOC entry 238 (class 1259 OID 16575)
-- Name: respostas_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.respostas_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.respostas_id_seq OWNER TO admin;

--
-- TOC entry 3573 (class 0 OID 0)
-- Dependencies: 238
-- Name: respostas_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.respostas_id_seq OWNED BY public.respostas.id;


--
-- TOC entry 218 (class 1259 OID 16386)
-- Name: usuarios; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.usuarios (
    id bigint NOT NULL,
    tipo_usuario smallint NOT NULL,
    nome character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    senha text NOT NULL,
    contato character varying(11),
    bio text,
    cpf character varying(11),
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    CONSTRAINT chk_usuarios_tipo_usuario CHECK ((tipo_usuario >= 1))
);


ALTER TABLE public.usuarios OWNER TO admin;

--
-- TOC entry 217 (class 1259 OID 16385)
-- Name: usuarios_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.usuarios_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.usuarios_id_seq OWNER TO admin;

--
-- TOC entry 3574 (class 0 OID 0)
-- Dependencies: 217
-- Name: usuarios_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.usuarios_id_seq OWNED BY public.usuarios.id;


--
-- TOC entry 3341 (class 2604 OID 16555)
-- Name: atribuicoes id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.atribuicoes ALTER COLUMN id SET DEFAULT nextval('public.atribuicoes_id_seq'::regclass);


--
-- TOC entry 3333 (class 2604 OID 16486)
-- Name: convites id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.convites ALTER COLUMN id SET DEFAULT nextval('public.convites_id_seq'::regclass);


--
-- TOC entry 3335 (class 2604 OID 16509)
-- Name: instrumentos id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.instrumentos ALTER COLUMN id SET DEFAULT nextval('public.instrumentos_id_seq'::regclass);


--
-- TOC entry 3330 (class 2604 OID 16470)
-- Name: notificacoes id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.notificacoes ALTER COLUMN id SET DEFAULT nextval('public.notificacoes_id_seq'::regclass);


--
-- TOC entry 3340 (class 2604 OID 16539)
-- Name: opcoes_escala id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.opcoes_escala ALTER COLUMN id SET DEFAULT nextval('public.opcoes_escala_id_seq'::regclass);


--
-- TOC entry 3327 (class 2604 OID 16421)
-- Name: pacientes id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.pacientes ALTER COLUMN id SET DEFAULT nextval('public.pacientes_id_seq'::regclass);


--
-- TOC entry 3338 (class 2604 OID 16522)
-- Name: perguntas id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.perguntas ALTER COLUMN id SET DEFAULT nextval('public.perguntas_id_seq'::regclass);


--
-- TOC entry 3326 (class 2604 OID 16404)
-- Name: profissionais id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.profissionais ALTER COLUMN id SET DEFAULT nextval('public.profissionais_id_seq'::regclass);


--
-- TOC entry 3328 (class 2604 OID 16451)
-- Name: registros_humor id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.registros_humor ALTER COLUMN id SET DEFAULT nextval('public.registros_humor_id_seq'::regclass);


--
-- TOC entry 3343 (class 2604 OID 16579)
-- Name: respostas id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.respostas ALTER COLUMN id SET DEFAULT nextval('public.respostas_id_seq'::regclass);


--
-- TOC entry 3325 (class 2604 OID 16389)
-- Name: usuarios id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.usuarios ALTER COLUMN id SET DEFAULT nextval('public.usuarios_id_seq'::regclass);


--
-- TOC entry 3392 (class 2606 OID 16560)
-- Name: atribuicoes atribuicoes_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.atribuicoes
    ADD CONSTRAINT atribuicoes_pkey PRIMARY KEY (id);


--
-- TOC entry 3375 (class 2606 OID 16491)
-- Name: convites convites_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.convites
    ADD CONSTRAINT convites_pkey PRIMARY KEY (id);


--
-- TOC entry 3382 (class 2606 OID 16515)
-- Name: instrumentos instrumentos_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.instrumentos
    ADD CONSTRAINT instrumentos_pkey PRIMARY KEY (id);


--
-- TOC entry 3373 (class 2606 OID 16476)
-- Name: notificacoes notificacoes_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.notificacoes
    ADD CONSTRAINT notificacoes_pkey PRIMARY KEY (id);


--
-- TOC entry 3390 (class 2606 OID 16543)
-- Name: opcoes_escala opcoes_escala_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.opcoes_escala
    ADD CONSTRAINT opcoes_escala_pkey PRIMARY KEY (id);


--
-- TOC entry 3365 (class 2606 OID 16423)
-- Name: pacientes pacientes_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.pacientes
    ADD CONSTRAINT pacientes_pkey PRIMARY KEY (id);


--
-- TOC entry 3386 (class 2606 OID 16527)
-- Name: perguntas perguntas_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.perguntas
    ADD CONSTRAINT perguntas_pkey PRIMARY KEY (id);


--
-- TOC entry 3358 (class 2606 OID 16406)
-- Name: profissionais profissionais_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.profissionais
    ADD CONSTRAINT profissionais_pkey PRIMARY KEY (id);


--
-- TOC entry 3369 (class 2606 OID 16436)
-- Name: profissional_paciente profissional_paciente_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.profissional_paciente
    ADD CONSTRAINT profissional_paciente_pkey PRIMARY KEY (paciente_id, profissional_id);


--
-- TOC entry 3371 (class 2606 OID 16460)
-- Name: registros_humor registros_humor_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.registros_humor
    ADD CONSTRAINT registros_humor_pkey PRIMARY KEY (id);


--
-- TOC entry 3400 (class 2606 OID 16583)
-- Name: respostas respostas_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.respostas
    ADD CONSTRAINT respostas_pkey PRIMARY KEY (id);


--
-- TOC entry 3378 (class 2606 OID 16493)
-- Name: convites uni_convites_token; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.convites
    ADD CONSTRAINT uni_convites_token UNIQUE (token);


--
-- TOC entry 3367 (class 2606 OID 16425)
-- Name: pacientes uni_pacientes_usuario_id; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.pacientes
    ADD CONSTRAINT uni_pacientes_usuario_id UNIQUE (usuario_id);


--
-- TOC entry 3360 (class 2606 OID 16410)
-- Name: profissionais uni_profissionais_registro_profissional; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.profissionais
    ADD CONSTRAINT uni_profissionais_registro_profissional UNIQUE (registro_profissional);


--
-- TOC entry 3362 (class 2606 OID 16408)
-- Name: profissionais uni_profissionais_usuario_id; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.profissionais
    ADD CONSTRAINT uni_profissionais_usuario_id UNIQUE (usuario_id);


--
-- TOC entry 3351 (class 2606 OID 16398)
-- Name: usuarios uni_usuarios_cpf; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.usuarios
    ADD CONSTRAINT uni_usuarios_cpf UNIQUE (cpf);


--
-- TOC entry 3353 (class 2606 OID 16396)
-- Name: usuarios uni_usuarios_email; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.usuarios
    ADD CONSTRAINT uni_usuarios_email UNIQUE (email);


--
-- TOC entry 3355 (class 2606 OID 16394)
-- Name: usuarios usuarios_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.usuarios
    ADD CONSTRAINT usuarios_pkey PRIMARY KEY (id);


--
-- TOC entry 3393 (class 1259 OID 16571)
-- Name: idx_atribuicoes_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_atribuicoes_deleted_at ON public.atribuicoes USING btree (deleted_at);


--
-- TOC entry 3394 (class 1259 OID 16573)
-- Name: idx_atribuicoes_instrumento_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_atribuicoes_instrumento_id ON public.atribuicoes USING btree (instrumento_id);


--
-- TOC entry 3395 (class 1259 OID 16574)
-- Name: idx_atribuicoes_paciente_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_atribuicoes_paciente_id ON public.atribuicoes USING btree (paciente_id);


--
-- TOC entry 3396 (class 1259 OID 16572)
-- Name: idx_atribuicoes_status; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_atribuicoes_status ON public.atribuicoes USING btree (status);


--
-- TOC entry 3376 (class 1259 OID 16504)
-- Name: idx_convites_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_convites_deleted_at ON public.convites USING btree (deleted_at);


--
-- TOC entry 3379 (class 1259 OID 16517)
-- Name: idx_instrumentos_codigo; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX idx_instrumentos_codigo ON public.instrumentos USING btree (codigo);


--
-- TOC entry 3380 (class 1259 OID 16516)
-- Name: idx_instrumentos_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_instrumentos_deleted_at ON public.instrumentos USING btree (deleted_at);


--
-- TOC entry 3387 (class 1259 OID 16549)
-- Name: idx_opcoes_escala_instrumento_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_opcoes_escala_instrumento_id ON public.opcoes_escala USING btree (instrumento_id);


--
-- TOC entry 3363 (class 1259 OID 16431)
-- Name: idx_pacientes_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_pacientes_deleted_at ON public.pacientes USING btree (deleted_at);


--
-- TOC entry 3383 (class 1259 OID 16533)
-- Name: idx_perguntas_instrumento_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_perguntas_instrumento_id ON public.perguntas USING btree (instrumento_id);


--
-- TOC entry 3356 (class 1259 OID 16416)
-- Name: idx_profissionais_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_profissionais_deleted_at ON public.profissionais USING btree (deleted_at);


--
-- TOC entry 3397 (class 1259 OID 16590)
-- Name: idx_respostas_atribuicao_id; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX idx_respostas_atribuicao_id ON public.respostas USING btree (atribuicao_id);


--
-- TOC entry 3398 (class 1259 OID 16589)
-- Name: idx_respostas_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_respostas_deleted_at ON public.respostas USING btree (deleted_at);


--
-- TOC entry 3349 (class 1259 OID 16399)
-- Name: idx_usuarios_deleted_at; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX idx_usuarios_deleted_at ON public.usuarios USING btree (deleted_at);


--
-- TOC entry 3388 (class 1259 OID 16550)
-- Name: index_composto_opcao_escala; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX index_composto_opcao_escala ON public.opcoes_escala USING btree (instrumento_id, valor, rotulo);


--
-- TOC entry 3384 (class 1259 OID 16534)
-- Name: index_composto_pergunta; Type: INDEX; Schema: public; Owner: admin
--

CREATE UNIQUE INDEX index_composto_pergunta ON public.perguntas USING btree (instrumento_id, ordem_item, dominio, conteudo);


--
-- TOC entry 3411 (class 2606 OID 16561)
-- Name: atribuicoes fk_atribuicoes_instrumento; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.atribuicoes
    ADD CONSTRAINT fk_atribuicoes_instrumento FOREIGN KEY (instrumento_id) REFERENCES public.instrumentos(id);


--
-- TOC entry 3412 (class 2606 OID 16566)
-- Name: atribuicoes fk_atribuicoes_paciente; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.atribuicoes
    ADD CONSTRAINT fk_atribuicoes_paciente FOREIGN KEY (paciente_id) REFERENCES public.pacientes(id);


--
-- TOC entry 3413 (class 2606 OID 16584)
-- Name: respostas fk_atribuicoes_resposta; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.respostas
    ADD CONSTRAINT fk_atribuicoes_resposta FOREIGN KEY (atribuicao_id) REFERENCES public.atribuicoes(id);


--
-- TOC entry 3407 (class 2606 OID 16499)
-- Name: convites fk_convites_paciente; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.convites
    ADD CONSTRAINT fk_convites_paciente FOREIGN KEY (paciente_id) REFERENCES public.pacientes(id) ON DELETE CASCADE;


--
-- TOC entry 3408 (class 2606 OID 16494)
-- Name: convites fk_convites_profissional; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.convites
    ADD CONSTRAINT fk_convites_profissional FOREIGN KEY (profissional_id) REFERENCES public.profissionais(id) ON DELETE CASCADE;


--
-- TOC entry 3410 (class 2606 OID 16544)
-- Name: opcoes_escala fk_instrumentos_opcoes_escala; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.opcoes_escala
    ADD CONSTRAINT fk_instrumentos_opcoes_escala FOREIGN KEY (instrumento_id) REFERENCES public.instrumentos(id) ON DELETE CASCADE;


--
-- TOC entry 3409 (class 2606 OID 16528)
-- Name: perguntas fk_instrumentos_perguntas; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.perguntas
    ADD CONSTRAINT fk_instrumentos_perguntas FOREIGN KEY (instrumento_id) REFERENCES public.instrumentos(id) ON DELETE CASCADE;


--
-- TOC entry 3406 (class 2606 OID 16477)
-- Name: notificacoes fk_notificacoes_usuario; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.notificacoes
    ADD CONSTRAINT fk_notificacoes_usuario FOREIGN KEY (usuario_id) REFERENCES public.usuarios(id) ON DELETE CASCADE;


--
-- TOC entry 3402 (class 2606 OID 16426)
-- Name: pacientes fk_pacientes_usuario; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.pacientes
    ADD CONSTRAINT fk_pacientes_usuario FOREIGN KEY (usuario_id) REFERENCES public.usuarios(id) ON DELETE CASCADE;


--
-- TOC entry 3401 (class 2606 OID 16411)
-- Name: profissionais fk_profissionais_usuario; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.profissionais
    ADD CONSTRAINT fk_profissionais_usuario FOREIGN KEY (usuario_id) REFERENCES public.usuarios(id) ON DELETE CASCADE;


--
-- TOC entry 3403 (class 2606 OID 16437)
-- Name: profissional_paciente fk_profissional_paciente_paciente; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.profissional_paciente
    ADD CONSTRAINT fk_profissional_paciente_paciente FOREIGN KEY (paciente_id) REFERENCES public.pacientes(id) ON DELETE CASCADE;


--
-- TOC entry 3404 (class 2606 OID 16442)
-- Name: profissional_paciente fk_profissional_paciente_profissional; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.profissional_paciente
    ADD CONSTRAINT fk_profissional_paciente_profissional FOREIGN KEY (profissional_id) REFERENCES public.profissionais(id) ON DELETE CASCADE;


--
-- TOC entry 3405 (class 2606 OID 16461)
-- Name: registros_humor fk_registros_humor_paciente; Type: FK CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.registros_humor
    ADD CONSTRAINT fk_registros_humor_paciente FOREIGN KEY (paciente_id) REFERENCES public.pacientes(id) ON DELETE CASCADE;


-- Completed on 2025-11-23 18:47:07 -03

--
-- PostgreSQL database dump complete
--

