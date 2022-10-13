INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/block/getByNum/@', '根据区块高度查询区块信息', 1, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/block/getByHash/@', '根据区块哈希查询区块信息', 1, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/block/maxBlock', '获取最大区块高度', 1, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/block/query', '区块的分页查询', 1, 'POST');


INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/tx/get/@', '根据交易hash获取交易详情', 2, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/tx/query', '交易的分页查询', 2, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/tx/getByAddr', '根据地址查询交易信息', 2, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/tx/amount', '获取历史交易总数量', 2, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/tx/getAmountAddr/@', '获取当前地址下交易总数量', 2, 'GET');


INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/contract/getByAddr/@', '通过地址查询合约信息', 3, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/contract/getTxAmount/@', '获取该合约地址下的交易数量', 3, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/contract/amount', '查询合约数量', 3, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/contract/query', '合约分页查询', 3, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/contract/token/query', 'token分页查询', 3, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/contract/holder', '该合约地址持有人信息', 3, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/contract/balance/top/holder', '该合约地址Top持有人详情信息', 3, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/contract/getToken', '根据token模糊查询合约地址', 3, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/contract/addedToken/info', '新增代币的信息', 3, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/contract/data/decode/@', '解析交易中的data数据', 3, 'GET');

INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/account/getByAddr/@', '通过address查询地址信息', 4, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/account/extra/info/@', '查询地址的附加信息', 4, 'GET');


INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/erc20/getByAddr/@', '通过address查询erc20信息', 5, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/erc20/getByContract/@', '通过合约地址查询erc20信息', 5, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/erc20/getByHash/@', '查询hash对应blockNum下面的erc20信息', 5, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/erc20/query', 'erc20信息分页查询', 5, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/erc20/getAmountAddr/@', '获取当前地址下erc20信息的总数量', 5, 'GET');


INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/relation/getByAddr/@', '根据地址获取relation树形图', 6, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/relation/getRoot', '获取所有Root地址', 6, 'GET');


INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/statistics/info/@', '根据地址获取统计信息', 7, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/statistics/tx', '根据地址获取交易统计信息', 7, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/*/statistics/holder/changes/@', '代币持有人数的历史变化', 7, 'GET');

INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/lpPool/query', 'lp_pool 分页查询', 9, 'POST');





INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/add', '增加标签', 10, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/get/@', '获取该用户下的所有标签', 10, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/get', '获取该用户对应地址的标签信息', 10, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/get/tags/@', '获取该地址下的所有标签', 10, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/remove', '删除某标签', 10, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/query', '分页查询标签', 10, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/query/tag', '通过标签查询', 10, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/group/query', '获取和当前用户下同一组下的所有不重复标签', 10, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/selectByProjectId/@', '根据项目ID获取标签', 10, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/classify/address', '根据地址获取分类标签（系统、个人、项目）', 10, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/user/personal/tag', '获取用户个人标签', 10, 'GET');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/label/query/personal/tag', '根据tag 查询个人和系统标签', 10, 'POST');

INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/address/get/tree/@', '根据地址查询gaaddress信息', 11, 'GET');


INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/project/add', '新增项目', 12, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/project/query', '分页查询项目', 12, 'POST');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/project/update', '更新项目', 12, 'PUT');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/project/delete/@', '删除项目', 12, 'DELETE');
INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/project/findById/@', '根据ID查询项目信息', 12, 'GET');


-- INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/casbin/role/@', '', '', 'GET');
-- INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/casbin/get', '', '', 'GET');
-- INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/casbin/permission/@', '', '', 'GET');
-- INSERT INTO "public"."api"("url", "name", "group", "ty") VALUES ('/v1/sixr/roles', '', '', 'GET');
