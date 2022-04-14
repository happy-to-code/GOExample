CREATE TABLE IF NOT EXISTS "public"."tx_history_stat" (
  "address_id" int8 NOT NULL,
  "contract_id" int8 NOT NULL,
  "token_type" int2,
  "in_amount" numeric(150),
  "in_txs" int4,
  "out_amount" numeric(150),
  "out_txs" int4,
  "stat_month" date,
  "balance" numeric(150),
  "updated_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT "tx_history_stat_pkey" PRIMARY KEY ("address_id", "contract_id","stat_month")
) PARTITION BY RANGE (stat_month)
;

COMMENT ON COLUMN "public"."tx_history_stat"."token_type" IS '币种，
0 - 代币；
1 - 主币';

COMMENT ON COLUMN "public"."tx_history_stat"."in_amount" IS '转入的总金额';

COMMENT ON COLUMN "public"."tx_history_stat"."in_txs" IS '转入的交易总次数';

COMMENT ON COLUMN "public"."tx_history_stat"."out_amount" IS '转出的总金额';

COMMENT ON COLUMN "public"."tx_history_stat"."out_txs" IS '转出的交易总次数';

COMMENT ON COLUMN "public"."tx_history_stat"."stat_month" IS '上一次统计日期';

COMMENT ON COLUMN "public"."tx_history_stat"."balance" IS '当前余额';

COMMENT ON COLUMN "public"."tx_history_stat"."updated_at" IS '更新时间';


-- --------------------------------------------------------------------------------------
-- 2020年
CREATE TABLE public.tx_history_stat_2020 PARTITION OF public.tx_history_stat
    FOR VALUES FROM ('2020-01-01 00:00:00') TO ('2021-01-01 00:00:00')
    PARTITION BY RANGE (stat_month);
-- 2020年-月
CREATE TABLE public.tx_history_stat_202001 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-01-01 00:00:00') TO ('2020-02-01 00:00:00');
CREATE TABLE public.tx_history_stat_202002 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-02-01 00:00:00') TO ('2020-03-01 00:00:00');
CREATE TABLE public.tx_history_stat_202003 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-03-01 00:00:00') TO ('2020-04-01 00:00:00');
CREATE TABLE public.tx_history_stat_202004 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-04-01 00:00:00') TO ('2020-05-01 00:00:00');
CREATE TABLE public.tx_history_stat_202005 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-05-01 00:00:00') TO ('2020-06-01 00:00:00');
CREATE TABLE public.tx_history_stat_202006 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-06-01 00:00:00') TO ('2020-07-01 00:00:00');
CREATE TABLE public.tx_history_stat_202007 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-07-01 00:00:00') TO ('2020-08-01 00:00:00');
CREATE TABLE public.tx_history_stat_202008 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-08-01 00:00:00') TO ('2020-09-01 00:00:00');
CREATE TABLE public.tx_history_stat_202009 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-09-01 00:00:00') TO ('2020-10-01 00:00:00');
CREATE TABLE public.tx_history_stat_202010 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-10-01 00:00:00') TO ('2020-11-01 00:00:00');
CREATE TABLE public.tx_history_stat_202011 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-11-01 00:00:00') TO ('2020-12-01 00:00:00');
CREATE TABLE public.tx_history_stat_202012 PARTITION OF public.tx_history_stat_2020
    FOR VALUES FROM ('2020-12-01 00:00:00') TO ('2021-01-01 00:00:00');


-- --------------------------------------------------------------------------------------
-- 2021年
CREATE TABLE public.tx_history_stat_2021 PARTITION OF public.tx_history_stat
    FOR VALUES FROM ('2021-01-01 00:00:00') TO ('2022-01-01 00:00:00')
    PARTITION BY RANGE (stat_month);
-- 2021年-月
CREATE TABLE public.tx_history_stat_202101 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-01-01 00:00:00') TO ('2021-02-01 00:00:00');
CREATE TABLE public.tx_history_stat_202102 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-02-01 00:00:00') TO ('2021-03-01 00:00:00');
CREATE TABLE public.tx_history_stat_202103 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-03-01 00:00:00') TO ('2021-04-01 00:00:00');
CREATE TABLE public.tx_history_stat_202104 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-04-01 00:00:00') TO ('2021-05-01 00:00:00');
CREATE TABLE public.tx_history_stat_202105 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-05-01 00:00:00') TO ('2021-06-01 00:00:00');
CREATE TABLE public.tx_history_stat_202106 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-06-01 00:00:00') TO ('2021-07-01 00:00:00');
CREATE TABLE public.tx_history_stat_202107 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-07-01 00:00:00') TO ('2021-08-01 00:00:00');
CREATE TABLE public.tx_history_stat_202108 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-08-01 00:00:00') TO ('2021-09-01 00:00:00');
CREATE TABLE public.tx_history_stat_202109 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-09-01 00:00:00') TO ('2021-10-01 00:00:00');
CREATE TABLE public.tx_history_stat_202110 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-10-01 00:00:00') TO ('2021-11-01 00:00:00');
CREATE TABLE public.tx_history_stat_202111 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-11-01 00:00:00') TO ('2021-12-01 00:00:00');
CREATE TABLE public.tx_history_stat_202112 PARTITION OF public.tx_history_stat_2021
    FOR VALUES FROM ('2021-12-01 00:00:00') TO ('2022-01-01 00:00:00');


-- --------------------------------------------------------------------------------------
-- 2022年
CREATE TABLE public.tx_history_stat_2022 PARTITION OF public.tx_history_stat
    FOR VALUES FROM ('2022-01-01 00:00:00') TO ('2023-01-01 00:00:00')
    PARTITION BY RANGE (stat_month);
-- 2022年-月
CREATE TABLE public.tx_history_stat_202201 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-01-01 00:00:00') TO ('2022-02-01 00:00:00');
CREATE TABLE public.tx_history_stat_202202 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-02-01 00:00:00') TO ('2022-03-01 00:00:00');
CREATE TABLE public.tx_history_stat_202203 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-03-01 00:00:00') TO ('2022-04-01 00:00:00');
CREATE TABLE public.tx_history_stat_202204 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-04-01 00:00:00') TO ('2022-05-01 00:00:00');
CREATE TABLE public.tx_history_stat_202205 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-05-01 00:00:00') TO ('2022-06-01 00:00:00');
CREATE TABLE public.tx_history_stat_202206 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-06-01 00:00:00') TO ('2022-07-01 00:00:00');
CREATE TABLE public.tx_history_stat_202207 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-07-01 00:00:00') TO ('2022-08-01 00:00:00');
CREATE TABLE public.tx_history_stat_202208 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-08-01 00:00:00') TO ('2022-09-01 00:00:00');
CREATE TABLE public.tx_history_stat_202209 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-09-01 00:00:00') TO ('2022-10-01 00:00:00');
CREATE TABLE public.tx_history_stat_202210 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-10-01 00:00:00') TO ('2022-11-01 00:00:00');
CREATE TABLE public.tx_history_stat_202211 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-11-01 00:00:00') TO ('2022-12-01 00:00:00');
CREATE TABLE public.tx_history_stat_202212 PARTITION OF public.tx_history_stat_2022
    FOR VALUES FROM ('2022-12-01 00:00:00') TO ('2023-01-01 00:00:00');



-- --------------------------------------------------------------------------------------
-- 2023年
CREATE TABLE public.tx_history_stat_2023 PARTITION OF public.tx_history_stat
    FOR VALUES FROM ('2023-01-01 00:00:00') TO ('2024-01-01 00:00:00')
    PARTITION BY RANGE (stat_month);
-- 2023年-月
CREATE TABLE public.tx_history_stat_202301 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-01-01 00:00:00') TO ('2023-02-01 00:00:00');
CREATE TABLE public.tx_history_stat_202302 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-02-01 00:00:00') TO ('2023-03-01 00:00:00');
CREATE TABLE public.tx_history_stat_202303 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-03-01 00:00:00') TO ('2023-04-01 00:00:00');
CREATE TABLE public.tx_history_stat_202304 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-04-01 00:00:00') TO ('2023-05-01 00:00:00');
CREATE TABLE public.tx_history_stat_202305 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-05-01 00:00:00') TO ('2023-06-01 00:00:00');
CREATE TABLE public.tx_history_stat_202306 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-06-01 00:00:00') TO ('2023-07-01 00:00:00');
CREATE TABLE public.tx_history_stat_202307 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-07-01 00:00:00') TO ('2023-08-01 00:00:00');
CREATE TABLE public.tx_history_stat_202308 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-08-01 00:00:00') TO ('2023-09-01 00:00:00');
CREATE TABLE public.tx_history_stat_202309 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-09-01 00:00:00') TO ('2023-10-01 00:00:00');
CREATE TABLE public.tx_history_stat_202310 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-10-01 00:00:00') TO ('2023-11-01 00:00:00');
CREATE TABLE public.tx_history_stat_202311 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-11-01 00:00:00') TO ('2023-12-01 00:00:00');
CREATE TABLE public.tx_history_stat_202312 PARTITION OF public.tx_history_stat_2023
    FOR VALUES FROM ('2023-12-01 00:00:00') TO ('2024-01-01 00:00:00');
