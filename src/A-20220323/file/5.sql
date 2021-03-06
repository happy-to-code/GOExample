INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/sixr/casbin/get', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/sixr/casbin/permission/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/block/getByNum/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/block/getByHash/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/block/maxBlock', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/block/query', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/tx/get/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/tx/query', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/tx/getByAddr', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/tx/amount', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/tx/getAmountAddr/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/contract/getByAddr/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/contract/getTxAmount/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/contract/amount', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/contract/query', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/contract/token/query', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/contract/holder', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/contract/balance/top/holder', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/contract/getToken', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/contract/addedToken/info', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/contract/data/decode/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/account/getByAddr/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/erc20/getByHash/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/erc20/query', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/erc20/getAmountAddr/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/relation/getByAddr/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/relation/getRoot', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/statistics/info/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/statistics/tx', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/*/statistics/holder/changes/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/sixr/lpPool/query', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/sixr/label/add', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/sixr/label/get/@', 'GET', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/sixr/label/get', 'POST', '', '', '');
INSERT INTO "public"."casbin_rule" ("p_type", "v0", "v1", "v2", "v3", "v4", "v5")
VALUES ('p', 'developer', '/v1/sixr/label/get/tags/@', 'GET', '', '', '');

0x5c0a02ccbc5eb09eef6773a161c6ccbe746fd29b