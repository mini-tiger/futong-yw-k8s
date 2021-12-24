drop database if exists ftk8s;
create database ftk8s default character set utf8mb4;
use ftk8s;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `p_type` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES ('g', 'FTUSER-3178db1353864cdebe0402cce615ad29', 'role1', 'admin', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'FTUSER-b6b5fd90dd724f70971c86ac8986ea5d', 'role1', 'admin', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'FTUSER-ec197677ac26430da65cb552642a612b', 'role1', 'admin', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'FTGROUP-1c621631428048aba147a73f0a4044e3', 'role1', 'admin', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'FTGROUP-61eedad88efa4fc2a9df83b7111ebd8a', 'role1', 'admin', '', '', '');
INSERT INTO `casbin_rule` VALUES ('g', 'FTGROUP-83f060962723409b9121f43e0555bcff', 'role1', 'admin', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/configmap', 'post', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/deployment', 'post', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/deployment-ui', 'post', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/namespace', 'post', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/service', 'post', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/configmap', 'delete', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/deployment', 'delete', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/namespace', 'delete', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/service', 'delete', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/deploy-template', 'post', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/configmaps', 'get', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/deployments', 'get', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/namespaces', 'get', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/services', 'get', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/configmap', 'get', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/deployment', 'get', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/namespace', 'get', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/service', 'get', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/configmap', 'put', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/deployment', 'put', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/deployment-ui', 'put', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/namespace', 'put', '', '');
INSERT INTO `casbin_rule` VALUES ('p', 'role1', 'admin', '/api/resource/service', 'put', '', '');
COMMIT;

-- ----------------------------
-- Table structure for cluster
-- ----------------------------
DROP TABLE IF EXISTS `cluster`;
CREATE TABLE `cluster` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键,集群ID',
  `tenant_id` varchar(150) NOT NULL COMMENT '租户ID',
  `cluster_account` varchar(150) NOT NULL COMMENT '集群账号',
  `cluster_name` varchar(150) NOT NULL COMMENT '集群名称',
  `cluster_api` varchar(150) NOT NULL COMMENT '集群API',
  `k8s_config` longtext NOT NULL COMMENT '集群秘钥',
  `description` text COMMENT '集群介绍',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_tenant_id_cluster_account` (`tenant_id`,`cluster_account`),
  UNIQUE KEY `ui_tenant_id_cluster_api` (`tenant_id`,`cluster_api`),
  CONSTRAINT `fk_cluster_cluster_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of cluster
-- ----------------------------
BEGIN;
INSERT INTO `cluster` VALUES (1, 'admin', 'aliyun', '阿里云k8s集群', 'https://39.97.223.173:6443', '{\"apiVersion\":\"v1\",\"clusters\":[{\"cluster\":{\"server\":\"https://39.97.223.173:6443\",\"certificate-authority-data\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURHakNDQWdLZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREErTVNjd0ZBWURWUVFLRXcxaGJHbGkKWVdKaElHTnNiM1ZrTUE4R0ExVUVDaE1JYUdGdVozcG9iM1V4RXpBUkJnTlZCQU1UQ210MVltVnlibVYwWlhNdwpIaGNOTWpBeE1ETXdNREkwT0RJeFdoY05NekF4TURJNE1ESTBPREl4V2pBK01TY3dGQVlEVlFRS0V3MWhiR2xpCllXSmhJR05zYjNWa01BOEdBMVVFQ2hNSWFHRnVaM3BvYjNVeEV6QVJCZ05WQkFNVENtdDFZbVZ5Ym1WMFpYTXcKZ2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRREkyZXcwZjRlVTBOd0Z1V3dUU3NscwpzNEs3Ym1mRWlnZnNzMENWdDRMR2FvbEJYZHlob0djWmV2VUFQWHhrYmxJSmVOcTBNditKSG5Ea0I4TjBoOTlTCjhpZVE0TUM2YkI4anJsb3pNcVowaHJtczdYODNDcVZTRVVhamlob1dBTTI2V1ovZFpqUjNBNHRWSWFwK3JUcTYKZUdQaks3Y2k5eEU1NS9rN2ErSWVldDBxOEpMaDdabDBtRlFjVHBhL2E0WjhUcTdGTWVkZmpqc1JYK1JxbVByeQp2bTNtVVFoNXNCN2VlcmM3d1IxKzYxMTNJTHcrakhDRHBqZ3pFcGRlV2dOZWQyL0RpSWhUYXFPY0ZiSlpnU1lXCkl5dmttbUxaZkxIWm5vQUhrd09pQlkxK1ZZZ25WUUVxeHhKek52ckdjcHN6cVI3bHZ0MndUdTltNWFBc0M3ZjEKQWdNQkFBR2pJekFoTUE0R0ExVWREd0VCL3dRRUF3SUNwREFQQmdOVkhSTUJBZjhFQlRBREFRSC9NQTBHQ1NxRwpTSWIzRFFFQkN3VUFBNElCQVFBYjBrZG9IMjlzQ2FsUFE5TWxPaUI4OEQ5NFVJWFJ4azNFUkFPT2J3eUhybi92Cm5SMENXb2RTTVo0Rng2UnNjQjRGVUdzaXVuZ0dhOEdiTWR0ckhjTWZaY2pFV21EMEk1NGVIaENOOTUrR3ZIcy8KdU5UY2w5TVY1aStVOGFObzNKaEFsUVd5WW9Ic0dhT2l0SGI2WGJFMzJDajVTUjNQMkpYTHdNSjc2YmZ4eC9WVApYdnFCYW5TblcvZzBrR2lRc2xDWFlwOHZZR0VtYnpVL0plaEovTFN0MTJvZC95YktpVjdzQUVuS0NGcks2VzFWCk9TeTdSOW82T3RXWCt0ekpDUWZ6aXQrZUlVYUJTczFZc2ZoYU1XN3VKeUdlOFo5WlY5WFZmdDQ4TCtyLzdNelMKZ1ZSUkpYVFBoTUIyV010endRaHhvZ3EvcStqWmNtT2JVNDN5N1RzZQotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\"},\"name\":\"kubernetes\"}],\"contexts\":[{\"context\":{\"cluster\":\"kubernetes\",\"user\":\"249201057743504030\"},\"name\":\"249201057743504030-ce3a4ec6eff4e451e89e2c010794697b7\"}],\"current-context\":\"249201057743504030-ce3a4ec6eff4e451e89e2c010794697b7\",\"kind\":\"Config\",\"preferences\":{},\"users\":[{\"name\":\"249201057743504030\",\"user\":{\"client-certificate-data\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURYRENDQXNXZ0F3SUJBZ0lERzFKck1BMEdDU3FHU0liM0RRRUJDd1VBTUdveEtqQW9CZ05WQkFvVElXTmwKTTJFMFpXTTJaV1ptTkdVME5URmxPRGxsTW1Nd01UQTNPVFEyT1RkaU56RVFNQTRHQTFVRUN4TUhaR1ZtWVhWcwpkREVxTUNnR0ExVUVBeE1oWTJVellUUmxZelpsWm1ZMFpUUTFNV1U0T1dVeVl6QXhNRGM1TkRZNU4ySTNNQjRYCkRUSXdNVEF6TURBeU5Ea3dNRm9YRFRJek1UQXpNREF5TlRRd01sb3dTakVWTUJNR0ExVUVDaE1NYzNsemRHVnQKT25WelpYSnpNUWt3QndZRFZRUUxFd0F4SmpBa0JnTlZCQU1USFRJME9USXdNVEExTnpjME16VXdOREF6TUMweApOakEwTURJMk5EUXlNSUlCSWpBTkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQTJOMlBvQVQ2CldnV0RmZGUyUHRuU3dpS25mY1pjSEpzOGxya0JlU29wQVlsWERrTHIycGJHOTVsR3Y0eHdORXNUenk1OU5yY00KV1lTUGprNlFod3FHRW5mWmtyN0ZJZ2RJcGtQY3dDRVVrS0pBd2lGbjUybCtlZUhLNlNOQnN6Rld6eHVJMUQwSwpWdW5mVXNsVDhJeXVSOEt6S2lWN0VyMzRXc29oejJ3R2tpMU9KS040MnN1RXlqZHZDVkFsajhWdHVvVlJjb3F0CkJCcGhlRXVtbnlrbXd1QzI3L0o2UkVkR29sTTlxYkhrTk8vckhTelhPaHM0OFEyNXBNZG0vQTJXQnliTXZRaUsKVi9Pd3A5MWtUVEl5TWdwVms2Q1h6bUI4cVhjeGMxeEZ3Y1Zad2FScmFaY1lnakZNOUJ1QjRKTlI4RE1XOXhhOQpCTi9lcjNXcUNsMEhkd0lEQVFBQm80R3JNSUdvTUE0R0ExVWREd0VCL3dRRUF3SUhnREFUQmdOVkhTVUVEREFLCkJnZ3JCZ0VGQlFjREFqQU1CZ05WSFJNQkFmOEVBakFBTUR3R0NDc0dBUVVGQndFQkJEQXdMakFzQmdnckJnRUYKQlFjd0FZWWdhSFIwY0RvdkwyTmxjblJ6TG1GamN5NWhiR2w1ZFc0dVkyOXRMMjlqYzNBd05RWURWUjBmQkM0dwpMREFxb0NpZ0pvWWthSFIwY0RvdkwyTmxjblJ6TG1GamN5NWhiR2w1ZFc0dVkyOXRMM0p2YjNRdVkzSnNNQTBHCkNTcUdTSWIzRFFFQkN3VUFBNEdCQUozLzR3SWxWYVptd3BmcVAvYTlpMlJKSFpReDd5L2k1R1E3eElwKzYxclQKYkVPSmRLamlRdTdiK3d5b0FFOWZ5VVdNenFyZTVJMlFraUI5dTE1bzJRcWx5ZERQZFFXcVpRNHJqR3dpcVNiTwpUbmF1djVsZ0kyL0lZYzVQdmQ4THdxVDZHZjVpSG5vK3BMR0wvTFJ5aisxcmNlVzJkWkZGdDJlR3FuNE0rKzlZCi0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0KLS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUMvekNDQW1pZ0F3SUJBZ0lERzFKV01BMEdDU3FHU0liM0RRRUJDd1VBTUdJeEN6QUpCZ05WQkFZVEFrTk8KTVJFd0R3WURWUVFJREFoYWFHVkthV0Z1WnpFUk1BOEdBMVVFQnd3SVNHRnVaMXBvYjNVeEVEQU9CZ05WQkFvTQpCMEZzYVdKaFltRXhEREFLQmdOVkJBc01BMEZEVXpFTk1Bc0dBMVVFQXd3RWNtOXZkREFlRncweU1ERXdNekF3Ck1qUXhNREJhRncwME1ERXdNalV3TWpRMk16WmFNR294S2pBb0JnTlZCQW9USVdObE0yRTBaV00yWldabU5HVTAKTlRGbE9EbGxNbU13TVRBM09UUTJPVGRpTnpFUU1BNEdBMVVFQ3hNSFpHVm1ZWFZzZERFcU1DZ0dBMVVFQXhNaApZMlV6WVRSbFl6WmxabVkwWlRRMU1XVTRPV1V5WXpBeE1EYzVORFk1TjJJM01JR2ZNQTBHQ1NxR1NJYjNEUUVCCkFRVUFBNEdOQURDQmlRS0JnUUMwSGx2TkNGUXYrb1pzWi9BcmZwK2dqdWV5NkRyNXlibXNSTHJzNlJCdFVRbC8KZTZER1U3YzF5VGJET1p6QTNRN0RJVVc1bU1vaUNYRk5ZOXFnZDllK1llZFpjaVNuSUdDN0w0d2xSb2JnR0tXdQpqM2ppMFlMVTA2dWtBNkdkbTlqWFU2RVJqSWRZanJHMVdzS2tBRnMwYVN6LzRBYXhPSHdmQy9jenAyMVRlUUlECkFRQUJvNEc2TUlHM01BNEdBMVVkRHdFQi93UUVBd0lDckRBUEJnTlZIUk1CQWY4RUJUQURBUUgvTUI4R0ExVWQKSXdRWU1CYUFGSVZhLzkwanpTVnZXRUZ2bm0xRk9adFlmWFgvTUR3R0NDc0dBUVVGQndFQkJEQXdMakFzQmdncgpCZ0VGQlFjd0FZWWdhSFIwY0RvdkwyTmxjblJ6TG1GamN5NWhiR2w1ZFc0dVkyOXRMMjlqYzNBd05RWURWUjBmCkJDNHdMREFxb0NpZ0pvWWthSFIwY0RvdkwyTmxjblJ6TG1GamN5NWhiR2w1ZFc0dVkyOXRMM0p2YjNRdVkzSnMKTUEwR0NTcUdTSWIzRFFFQkN3VUFBNEdCQUNpZUY4UzVvWEt0UzVOY2hvN0dEdjV4N2ZoT2lQQnJYTytaL29qUApHR1VpZU41YXM2aU1tNEN4MjFQZkxiQVJ6a1pkWWdCcklwTk5LVkVtYmo4ZTVGRmc1aVZnMGxRd05rNnlnWHBMClAwbW1RcmM3Sk5TL2NjRVA5UXcyNW84MFNSdXdmd0RXQ1FOelI1cFVPTmFBRDhxbEduUVFXUDIwL2I5Z2lTRy8KUDJYMgotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0tCg==\",\"client-key-data\":\"LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBMk4yUG9BVDZXZ1dEZmRlMlB0blN3aUtuZmNaY0hKczhscmtCZVNvcEFZbFhEa0xyCjJwYkc5NWxHdjR4d05Fc1R6eTU5TnJjTVdZU1BqazZRaHdxR0VuZlprcjdGSWdkSXBrUGN3Q0VVa0tKQXdpRm4KNTJsK2VlSEs2U05Cc3pGV3p4dUkxRDBLVnVuZlVzbFQ4SXl1UjhLektpVjdFcjM0V3NvaHoyd0draTFPSktONAoyc3VFeWpkdkNWQWxqOFZ0dW9WUmNvcXRCQnBoZUV1bW55a213dUMyNy9KNlJFZEdvbE05cWJIa05PL3JIU3pYCk9oczQ4UTI1cE1kbS9BMldCeWJNdlFpS1YvT3dwOTFrVFRJeU1ncFZrNkNYem1COHFYY3hjMXhGd2NWWndhUnIKYVpjWWdqRk05QnVCNEpOUjhETVc5eGE5Qk4vZXIzV3FDbDBIZHdJREFRQUJBb0lCQUM3Yk1EMFMxa1M5RER3VQpiM3dFOUZTZHlES1V1VEkxR0ZJNGh1ajNBd1VoOTMxTldFaFNhNHJ6d3lWLzRuNXArazI1YmJSMHVHWmZEZVZoCmREaTVQVjZnSnBKZVJabWttVDNUUzg3M1Zzb3BSSFN0WXhYTTVWYlFRbGM5RnVUd3RDRHJnaFRaVzNLTDlZU2QKbTFWT0VCblJKNFRqdEVQSVovQzEyN2hGVHVZbUhmcHBldGptN2N1VlQzVlFkNmQ0UnhhZnoyMDFyWlpOOEdkTwpJblJjbHFVbkJDNHRMelptenFZVVdiQ1hMV1l5amg3Qk9MQUdjN2hOUzVxOG15V3NxRnQvajhZOU9WZmh1L0hYClQvM3hSK09PSVBqWWhSTkJmckJ3N1VjVnAzVkZ4NkY4MHpYN2p3WnJ2OW1hVU52RzdqZ01wVk5aZmNxTHdhamoKQXJ3dGw0RUNnWUVBN0ZSb0NkbEpoTTRwMlJkcHRhRFNyL2d1bWhrbDc1UUI2QzYydktEeUJIYWlKOU8yQnlqUwpBSU1LM3Q2RHVMclpLSU1OdGdQSVJYMWdHRHFseTNHaWhBajUzTmtLNGxFQ0RZUXRwVXRicGhZUTBHVlJ4R29QCk9nWkdLMVd3OE9GTkxvOXVqQWgySWRoWVZKZWtSMUtsb0ZJZytBa3hueElGWHNHRmIyYWpJYk1DZ1lFQTZ1cHIKcFZoQlgyV1VMeEJHUFNzZDNpMCtCTzBkckx1c3Z1ZldISkhDRWx0VUZkQnNnOUhnTmZySDhNaVBzRXdZdFB1ZApzWWg5c2MrbGRlRXBCYWp3ZFJCQmozWFZ2c2xyV1UvaW42dm5PQ0tFM1VmaGhWM1Z0WERkd2JSUzE3ZHhrbFVoCmlmN00vR201bmJsczZsRlV4K2pMVEZnVHF3bUlhcnFuYWMzeitTMENnWUVBaTBwNGc0MkpNbmhjMi9KYndNeUkKUVdVeSttcG1ISjRNdmErQ3p3ekJlSFgvdngrZVF2d0JRb0g2RHoveFBSa0wwei9pL2V1ZXg2NU16QnNOQ3lydgowWWlFMUhFc1pCWEE4dng3OXRmQ3JkS0ZtSDZQUWdnTTczTXhPbXRvUGFGZFgxcjBtaDZHbWc0c2liZFRBU0txCk1pTVdBWHRSdnVMZFBXc3NYV3VPTE5jQ2dZQmllUUhHdmNncGhSc28zMW1TS3BES29ZeHQ1RGVjUU0rWTl2WDUKNDkvR2NpSTlRckU2VUsvNzhMUC9heE5RZzVXWHlDSENXY1RXMUlRM2Ric01kRlRYdllpOTZsYy9NbytkVUs5ZQozMkI5VjNyNmh6Y3lBTE9rNnpzS0I5enlNQ09iZzdRYmRJUFpDemlUdFRiQ3duVVhsNjRnd29yOEYzTG55bmFECkU2SkJ3UUtCZ1FDaS9GcGUzSms3UERLcmpiZ21qYUFnSHQyQ0JZZ08zQTBMSGJhTU5EcmhmSmE2elpRSUlYejAKK3ZWMzVzaEo5RWlnSC92bmt5K0lJNXJPd1pxZFRlSCtQZ1NJUThzSVc3OCtFNEZoWnc0QWFwR00zZU1sRnpXMApsSTFUbGNuWHRtWlZKL2p6SjhWMDRqVWtHclcySno3d1dZVk11UUJlZTFSRnNGT3hOLzBSWWc9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\"}}]}', '阿里云k8s集群', '2020-10-30 09:07:11.307', '2020-10-30 09:14:08.949');
INSERT INTO `cluster` VALUES (2, 'admin', 'k8s_demo1', 'k8s-示例1', 'https://172.16.53.124:6443', '{\"apiVersion\":\"v1\",\"clusters\":[{\"cluster\":{\"certificate-authority-data\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUN5akNDQWJLZ0F3SUJBZ0lCQURBTkJna3Foa2lHOXcwQkFRc0ZBREFWTVJNd0VRWURWUVFERXdwcmRXSmwKY201bGRHVnpNQ0FYRFRJd01UQXlNakV6TURJd05Gb1lEekl4TWpBd09USTRNVE13TWpBMFdqQVZNUk13RVFZRApWUVFERXdwcmRXSmxjbTVsZEdWek1JSUJJakFOQmdrcWhraUc5dzBCQVFFRkFBT0NBUThBTUlJQkNnS0NBUUVBCnFvdEVJYnZ1L0tyMUoyRC8zNkN5QUtMcFQ2V2wreE05ZjJuTno4VlpubFdrZHYvNUR4QksxWHh2SXg1MDh1bzUKd1d2ME5rdDFOME5QNnhKTTZCWFdjMDFpVTc5bWNWa1lUN2VVQkk5ZTFRTmVFOXZJSjd0NGIzSHl0aXN2NFZNNwpDVHNQWjdMUUx1TDBFVXFkZXBkZ29QaTlzb1NJZ0VCVjZaRzM2S0FxaEtGSVJaT0pyamtTa0ZGVkxRbXl3aHNwCkdDSmRWSGMwamdYNmoyNEZqVE9WOW1GOXQ3VWJzWnRMcTFBT240OTh5YXkwRmRNUE10QVp1QzFtUWpGeVY1NWIKMERGWXg2ZHBYVmszNkpXNzNDWEpiek1USlR3MGVjS3NkY1VvT01xVVBTa1hyYUdWSllDQmJLZS8yUkRTL3RKZQpabk5NQ1dueFlxVC9oTFp2ckt2SXV3SURBUUFCb3lNd0lUQU9CZ05WSFE4QkFmOEVCQU1DQXFRd0R3WURWUjBUCkFRSC9CQVV3QXdFQi96QU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFkZytnaEtrNTlFK2Y2RmlUMVVyNFZleDYKWm53a3IxYjNhUHNuRUlHQURjVVlOYkZua3ZBZzRiQzZFZWVjcXBEUHpSRllHdDN3czhhNFg2VXlDaVFteWpENgpNb2F1c1NuQkRUSzBaS2FkaG5hY2FCajlxNzNuU1dtTlBqVGFYL3lCWVhnT2FNS3ZvbzFUVlJrY2dZY04yUDFXCjYzYUE1RksveFgzVGJyNlgyMEpLdXE1RXcwNm0xb09PNG9kcmN0bWcwSTMrbzd6TkRzVkoxaEszaGVjekhlR2cKVE5LcjVSNWxjYVFwTlhiNU82c1JQSDdCcXpHQktMRGlkRGFnTi9MQkhhcUdBN2VnYVV1eFhJZkxYeHFpT3ZTWApVWWNLUkI0MEw2RlRXcndaVWZESUMzeEZaa2tuem4rZCtlWEdoWXlkKzc4K2VYNlNqSVZPL1Jxb1dxQ3Fqdz09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\",\"server\":\"https://172.16.53.124:6443\"},\"name\":\"kubernetes\"}],\"contexts\":[{\"context\":{\"cluster\":\"kubernetes\",\"user\":\"kubernetes-admin\"},\"name\":\"kubernetes-admin@kubernetes\"}],\"current-context\":\"kubernetes-admin@kubernetes\",\"kind\":\"Config\",\"preferences\":{},\"users\":[{\"name\":\"kubernetes-admin\",\"user\":{\"client-certificate-data\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM5RENDQWR5Z0F3SUJBZ0lJZEFTWmVWdHBjTkF3RFFZSktvWklodmNOQVFFTEJRQXdGVEVUTUJFR0ExVUUKQXhNS2EzVmlaWEp1WlhSbGN6QWdGdzB5TURFd01qSXhNekF5TURSYUdBOHlNVEl3TURreU9ERXpNRE13TkZvdwpOREVYTUJVR0ExVUVDaE1PYzNsemRHVnRPbTFoYzNSbGNuTXhHVEFYQmdOVkJBTVRFR3QxWW1WeWJtVjBaWE10CllXUnRhVzR3Z2dFaU1BMEdDU3FHU0liM0RRRUJBUVVBQTRJQkR3QXdnZ0VLQW9JQkFRREl3NnE1WXlpMFRSMWkKMXVQNnNPQWhPQzB4SUxkWTBTb2IrTWhBS0kySmRpY25US0p6czE5ek53cGJRVjlFOWtVNFU2YmV0L0VFSFpuRwpiRkxqMGRjUTlPYTlGMjRicVlSblNiR1RqWTE1VDhqSTV0aU1tbnFxa3I0SEg1b3hGU3Y5QVhhR1ZIZDlpaEgvCmZqNlFPOXlvTm1sY3h1VGM3OHcxQzVBeE9FOHNFVVVaelg5RmJyUVlxUVFxd25sNStybUN0WmZjN09Kb1h2S0sKY2ZJNEtVTWRCN1VoYi8xcTQ5K0dhUVVvaUJrLzVaUUlJaTc4OHkzYVF5bTZVQnVKRTNzV3J5Q2dYSTJjcWJVYQpJc2ZidVRuNzBVenVmZnNjTlVEdEtyaDBuTjFOWUwxQk5NTGIwTDN0ZjhOWmg1Mk5UL20wQjhqT3RnS3FBUnA0CnE3WXlhei8vQWdNQkFBR2pKekFsTUE0R0ExVWREd0VCL3dRRUF3SUZvREFUQmdOVkhTVUVEREFLQmdnckJnRUYKQlFjREFqQU5CZ2txaGtpRzl3MEJBUXNGQUFPQ0FRRUFGbWszQ0FJK2xZNG1INnZvUHg5UnpkM1FCSFU4aElZSQo5OFJES2RMTzNsY3laWVNHQm5iV2R3MURmSGIrWlNmNmNrbXhwWHZoWlFodGN2aW5jSmd6SnErTjl2S2UrR0tWClFHcGprdUdWMUJyUW1MVHdWbTQ2QzgxbmxzdHN3VXNxRDVwa3poc3J6VFlSWkEyL2g1Uy9rUXl0Nlo5YWR0WjgKdmxManU4TTRIelZUMjdJdloyUHVtWTllc3lCZmtNMXZXS0xaMUpYUk9Yd014MWx4cVFHYUVqU2hoRmg4TXI2aAp2cmFJcVFYcTllWGxNd3F1MzJLS1lwWVdLWFNFOW1YOGRwVXRTR2RWdjZkUUJKUlBvaWpDNm00dTBuMitSQ0NKClloN29tMDE0aTFueGwvbSthZGFVbFhzeDFqOFBjOVJQRHF2VEdRcTNWQkVqZDRtTDYyQnpNQT09Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\",\"client-key-data\":\"LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcEFJQkFBS0NBUUVBeU1PcXVXTW90RTBkWXRiaityRGdJVGd0TVNDM1dORXFHL2pJUUNpTmlYWW5KMHlpCmM3TmZjemNLVzBGZlJQWkZPRk9tM3JmeEJCMlp4bXhTNDlIWEVQVG12UmR1RzZtRVowbXhrNDJOZVUvSXlPYlkKakpwNnFwSytCeCthTVJVci9RRjJobFIzZllvUi8zNCtrRHZjcURacFhNYmszTy9NTlF1UU1UaFBMQkZGR2MxLwpSVzYwR0trRUtzSjVlZnE1Z3JXWDNPemlhRjd5aW5IeU9DbERIUWUxSVcvOWF1UGZobWtGS0lnWlArV1VDQ0l1Ci9QTXQya01wdWxBYmlSTjdGcThnb0Z5Tm5LbTFHaUxIMjdrNSs5Rk03bjM3SERWQTdTcTRkSnpkVFdDOVFUVEMKMjlDOTdYL0RXWWVkalUvNXRBZkl6cllDcWdFYWVLdTJNbXMvL3dJREFRQUJBb0lCQUFFQWNPNmgzd2NmUjJGQgozenRWL0poTjFuUGpUT2JsakVjOWM0cFdhWFpoSDRyanAvL1p5a1VoNWl4VVpDeE02a1dBclZsNUkzdTR4aGFtClhiZURTWVp0SW1XWkkxU0NBUVllNlFMcWR1VS9ENnBvOUhXbkk2dU1OZVNGTk5pLzJVdFc0WWVFRG1DUytzb0MKa0Q5Wi9SemR3S0xVM3psMi9OYmE2dXBEOUtrc25wdHpYN2xkbVNJbGE3djFwUFJMOWhGWnRBU3c1d1F4VTFOeQpEY0U2TnpHcFZWMlIxS1RqS0I1aFA2Q2kyV0JTc21UN2F2WXRDOWcvM2kwTzZybkJvMGNLOHc1MG1CUUF5NzltCkRCS1JnMURXNkF2WXQrY2tqdnJlZVAvdTVWT3k5dEUyOGIveXEwamNOOWQ0UzRhWi93RHBERTV6bHJCSllmcjUKR0V4K0lsa0NnWUVBeVh5N1E5Z1pORFh5cXk2dWc2dzd3bFo2S2xxN2VQV1Q4ZlltT1ArdTdYd3p5TC9vbUtvdwpuYURNMDloNlAyMEtwQ2RvWCt1aHR3MU1YL3k2N1BmS1c2WUhyLzNIdXYzanlvY0FHMVZPZnViSmN1YmNaYVlFCjhoUWVnZ2JGcW1McVFoVUpreTI0THZYYzl4ZHBvWTFVMDNnREpwd08xeDZFZW9MNmMyWjYvN1VDZ1lFQS94VGQKcElsRDRycGFEK0VwOHJma1FoVzFSSDJPelpUNHlTeW91cEQzUTdmbXNQVldzQ0VtVmRUaDNOby84ZlUrbFZEUwo1NmVmdkQ2MHNZZDkzV2VDUHdGTzViL1NRSW03N3g2bDAvaENIS3Bpd1N2VDRrckU5VVB5ZXdFYVVveEtQR2NTCktZQmNmSXY5RDlSb1hGL2R0WkxnYzUzOUxUVlFreWU4eW82SkNXTUNnWUFiWmlyUS93SHVaNmVvUDZEYnB3QTUKWFNrZnVWYmxEQUFpVnlhN2VZbUFJZk9veVBBSVVweHAwd1FIaXRpVTUyOGJBNERlQ2x6alY1dWJNZkw1Wm5qTwpIYlhONk9UUU9OWlJKQ3FQalBvdnA1S2RYV2Q3S3loaEF2dGpFeWl1RXVWb054UW1QNEZjWVhLNjV5UW1JK0gxCkg4ek40MHJoVmZVTjgrRzY3NlNxUVFLQmdRQ2hacGFJRWNLK3A2TnRBTkFHSUkxeHMwS3JQN2NvSmViMWhDaVAKUkVMd3VtQmlBQnNGL3pPK2c2RVdtWkowaXZVNmpaV2x0czYvMGYyTGgyd0F1QW9WeThJY1phK24zbjduNHNDUwp5emNwNURYd2ZNYnBITjViUXM4ZlBhZG1MQmFjK2Fyb1Q2dzFzbTVCb2VrVzZpSGpUamh2Yjd4TzZyazlJNXUxCm4xTGVlUUtCZ1FDVGNKS2VyclhJbmR2MGxLNkNjbGpUU1lraWZZSFVDZzMwRTV1ZFV3N3Z6Q2l1S0Z5ZXozSU0KZUtOeXJINFpJancyUk9vUnZjOWhZckpISnRRa3RQUnJRWXVoclE2L250YlltaDlmL2dtUXp1SUZ0MEFsa3BRbAp6ejRxZWpXYlFvOHZLTUtRbnZzUUFnTEsvVDc1R20yZWJla3ZlanJkR3FWMjhhdmJEUFdpblE9PQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\"}}]}', 'k8s-示例1的描述', '2020-10-27 07:47:16.744', '2020-10-30 07:11:31.356');
COMMIT;

-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` varchar(150) NOT NULL COMMENT '主键,用户组ID',
  `tenant_id` varchar(150) NOT NULL COMMENT '租户ID',
  `group_account` varchar(150) NOT NULL COMMENT '用户组账号',
  `group_name` varchar(150) NOT NULL COMMENT '用户组名称',
  `description` text COMMENT '用户组介绍',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_tenant_id_group_account` (`tenant_id`,`group_account`),
  CONSTRAINT `fk_group_group_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of group
-- ----------------------------
BEGIN;
INSERT INTO `group` VALUES ('FTGROUP-1c621631428048aba147a73f0a4044e3', 'admin', 'group3', '用户组3', '用户组3的描述信息', '2020-10-21 09:41:06.625', '2020-10-21 09:41:06.625');
INSERT INTO `group` VALUES ('FTGROUP-61eedad88efa4fc2a9df83b7111ebd8a', 'admin', 'group1', '用户组1', '用户组1的描述信息', '2020-10-21 09:40:56.716', '2020-10-21 09:40:56.716');
INSERT INTO `group` VALUES ('FTGROUP-83f060962723409b9121f43e0555bcff', 'admin', 'group2', '用户组2', '用户组2的描述信息', '2020-10-21 09:41:01.566', '2020-10-21 09:41:01.566');
COMMIT;

-- ----------------------------
-- Table structure for group_cluster_ass
-- ----------------------------
DROP TABLE IF EXISTS `group_cluster_ass`;
CREATE TABLE `group_cluster_ass` (
  `group_id` varchar(150) NOT NULL COMMENT '主键,用户组ID',
  `cluster_id` bigint NOT NULL COMMENT '主键,集群ID',
  PRIMARY KEY (`group_id`,`cluster_id`),
  KEY `fk_group_cluster_ass_cluster` (`cluster_id`),
  CONSTRAINT `fk_group_cluster_ass_cluster` FOREIGN KEY (`cluster_id`) REFERENCES `cluster` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_group_cluster_ass_group` FOREIGN KEY (`group_id`) REFERENCES `group` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of group_cluster_ass
-- ----------------------------
BEGIN;
INSERT INTO `group_cluster_ass` VALUES ('FTGROUP-61eedad88efa4fc2a9df83b7111ebd8a', 2);
COMMIT;

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
  `id` varchar(150) NOT NULL COMMENT '主键,权限ID',
  `permission_name` varchar(191) NOT NULL COMMENT '权限名称',
  `permission_url` varchar(191) NOT NULL COMMENT '权限URL',
  `permission_action` varchar(100) NOT NULL COMMENT '权限动作',
  `description` text COMMENT '权限介绍',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_permission_permission_name` (`permission_name`),
  UNIQUE KEY `ui_permission_url_permission_action` (`permission_url`,`permission_action`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of permission
-- ----------------------------
BEGIN;
INSERT INTO `permission` VALUES ('CreateConfigMap', '创建配置资源', '/api/resource/configmap', 'post', '创建配置资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreateCronJob', '创建定时任务资源', '/api/resource/cronjob', 'post', '创建定时任务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreateDaemonSet', '创建进程守护集资源', '/api/resource/daemonset', 'post', '创建进程守护集资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreateDeployment', '创建无状态资源', '/api/resource/deployment', 'post', '创建无状态资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreateIngress', '创建路由资源', '/api/resource/ingress', 'post', '创建路由资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreateJob', '创建任务资源', '/api/resource/job', 'post', '创建任务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreateNamespace', '创建命名空间', '/api/resource/namespace', 'post', '创建命名空间', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreatePV', '创建存储卷资源', '/api/resource/pv', 'post', '创建存储卷资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreatePVC', '创建存储声明资源', '/api/resource/pvc', 'post', '创建存储声明资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreateSecret', '创建保密字典资源', '/api/resource/secret', 'post', '创建保密字典资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreateService', '创建服务资源', '/api/resource/service', 'post', '创建服务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreateStatefulSet', '创建有状态资源', '/api/resource/statefulset', 'post', '创建有状态资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('CreateStorageClass', '创建存储类资源', '/api/resource/storageclass', 'post', '创建存储类资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeleteConfigMap', '删除配置资源', '/api/resource/configmap', 'delete', '删除配置资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeleteCronJob', '删除定时任务资源', '/api/resource/cronjob', 'delete', '删除定时任务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeleteDaemonSet', '删除进程守护集资源', '/api/resource/daemonset', 'delete', '删除进程守护集资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeleteDeployment', '删除无状态资源', '/api/resource/deployment', 'delete', '删除无状态资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeleteIngress', '删除路由资源', '/api/resource/ingress', 'delete', '删除路由资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeleteJob', '删除任务资源', '/api/resource/job', 'delete', '删除任务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeleteNamespace', '删除命名空间', '/api/resource/namespace', 'delete', '删除命名空间', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeletePV', '删除存储卷资源', '/api/resource/pv', 'delete', '删除存储卷资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeletePVC', '删除存储声明资源', '/api/resource/pvc', 'delete', '删除存储声明资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeleteSecret', '删除保密字典资源', '/api/resource/secret', 'delete', '删除保密字典资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeleteService', '删除服务资源', '/api/resource/service', 'delete', '删除服务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeleteStatefulSet', '删除有状态资源', '/api/resource/statefulset', 'delete', '删除有状态资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeleteStorageClass', '删除存储类资源', '/api/resource/storageclass', 'delete', '删除存储类资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('DeployTemplate', '部署模板', '/api/resource/deploy-template', 'post', '部署模板', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListConfigMap', '列表配置资源', '/api/resource/configmaps', 'get', '列表配置资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListCronJob', '列表定时任务资源', '/api/resource/cronjobs', 'get', '列表定时任务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListDaemonSet', '列表进程守护集资源', '/api/resource/daemonsets', 'get', '列表进程守护集资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListDeployment', '列表无状态资源', '/api/resource/deployments', 'get', '列表无状态资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListIngress', '列表路由资源', '/api/resource/ingresses', 'get', '列表路由资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListJob', '列表任务资源', '/api/resource/jobs', 'get', '列表任务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListNamespace', '列表命名空间', '/api/resource/namespaces', 'get', '列表命名空间', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListPod', '列表容器组资源', '/api/resource/pods', 'get', '列表容器组资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListPV', '列表存储卷资源', '/api/resource/pvs', 'get', '列表存储卷资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListPVC', '列表存储声明资源', '/api/resource/pvcs', 'get', '列表存储声明资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListSecret', '列表保密字典资源', '/api/resource/secrets', 'get', '列表保密字典资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListService', '列表服务资源', '/api/resource/services', 'get', '列表服务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListStatefulSet', '列表有状态资源', '/api/resource/statefulsets', 'get', '列表有状态资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ListStorageClass', '列表存储类资源', '/api/resource/storageclasses', 'get', '列表存储类资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadConfigMap', '获取配置资源', '/api/resource/configmap', 'get', '获取配置资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadCronJob', '获取定时任务资源', '/api/resource/cronjob', 'get', '获取定时任务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadDaemonSet', '获取进程守护集资源', '/api/resource/daemonset', 'get', '获取进程守护集资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadDeployment', '获取无状态资源', '/api/resource/deployment', 'get', '获取无状态资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadIngress', '获取路由资源', '/api/resource/ingress', 'get', '获取路由资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadJob', '获取任务资源', '/api/resource/job', 'get', '获取任务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadNamespace', '获取命名空间', '/api/resource/namespace', 'get', '获取命名空间', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadPod', '获取容器组资源', '/api/resource/pod', 'get', '获取容器组资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadPodLog', '读取容器组资源日志', '/api/resource/pod-log', 'get', '读取容器组资源日志', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadPV', '获取存储卷资源', '/api/resource/pv', 'get', '获取存储卷资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadPVC', '获取存储声明资源', '/api/resource/pvc', 'get', '获取存储声明资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadSecret', '获取保密字典资源', '/api/resource/secret', 'get', '获取保密字典资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadService', '获取服务资源', '/api/resource/service', 'get', '获取服务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadStatefulSet', '获取有状态资源', '/api/resource/statefulset', 'get', '获取有状态资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('ReadStorageClass', '获取存储类资源', '/api/resource/storageclass', 'get', '获取存储类资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdateConfigMap', '更新配置资源', '/api/resource/configmap', 'put', '更新配置资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdateCronJob', '更新定时任务资源', '/api/resource/cronjob', 'put', '更新定时任务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdateDaemonSet', '更新进程守护集资源', '/api/resource/daemonset', 'put', '更新进程守护集资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdateDeployment', '更新无状态资源', '/api/resource/deployment', 'put', '更新无状态资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdateIngress', '更新路由资源', '/api/resource/ingress', 'put', '更新路由资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdateJob', '更新任务资源', '/api/resource/job', 'put', '更新任务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdateNamespace', '更新命名空间', '/api/resource/namespace', 'put', '更新命名空间', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdatePV', '更新存储卷资源', '/api/resource/pv', 'put', '更新存储卷资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdatePVC', '更新存储声明资源', '/api/resource/pvc', 'put', '更新存储声明资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdateSecret', '更新保密字典资源', '/api/resource/secret', 'put', '更新保密字典资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdateService', '更新服务资源', '/api/resource/service', 'put', '更新服务资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdateStatefulSet', '更新有状态资源', '/api/resource/statefulset', 'put', '更新有状态资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
INSERT INTO `permission` VALUES ('UpdateStorageClass', '更新存储类资源', '/api/resource/storageclass', 'put', '更新存储类资源', '2020-11-13 10:01:22.385', '2020-11-13 10:01:22.385');
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键,角色ID',
  `tenant_id` varchar(150) NOT NULL COMMENT '租户ID',
  `role_account` varchar(150) NOT NULL COMMENT '角色账号',
  `role_name` varchar(150) NOT NULL COMMENT '角色名称',
  `description` text COMMENT '角色介绍',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_tenant_id_role_account` (`tenant_id`,`role_account`),
  CONSTRAINT `fk_role_role_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, 'admin', 'role1', '角色1', '所有菜单和数据权限', '2020-10-21 09:43:14.996', '2020-10-23 06:39:20.385');
COMMIT;

-- ----------------------------
-- Table structure for role_webperm_ass
-- ----------------------------
DROP TABLE IF EXISTS `role_webperm_ass`;
CREATE TABLE `role_webperm_ass` (
  `role_id` bigint NOT NULL COMMENT '主键,角色ID',
  `webperm_id` varchar(150) NOT NULL COMMENT '主键,前端权限ID',
  PRIMARY KEY (`role_id`,`webperm_id`),
  KEY `fk_role_webperm_ass_webperm` (`webperm_id`),
  CONSTRAINT `fk_role_webperm_ass_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_role_webperm_ass_webperm` FOREIGN KEY (`webperm_id`) REFERENCES `webperm` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of role_webperm_ass
-- ----------------------------
BEGIN;
INSERT INTO `role_webperm_ass` VALUES (1, '1');
INSERT INTO `role_webperm_ass` VALUES (1, '1_1');
INSERT INTO `role_webperm_ass` VALUES (1, '2c1bd867a4094d999afca43934fd1084');
INSERT INTO `role_webperm_ass` VALUES (1, '6c9c9fce4e9d487ebabad5da63b4aa94');
INSERT INTO `role_webperm_ass` VALUES (1, '720fb65c7c324f4d8fe724f839128c08');
INSERT INTO `role_webperm_ass` VALUES (1, '787bbc14376c48b1a58f9c73978f68af');
INSERT INTO `role_webperm_ass` VALUES (1, '8b02c88421c243f6a30ae415593c28bd');
INSERT INTO `role_webperm_ass` VALUES (1, '97a9799b474847549b7ee60d8440f69b');
INSERT INTO `role_webperm_ass` VALUES (1, 'a18d07bb29c84ad5befb20353f8dfa08');
INSERT INTO `role_webperm_ass` VALUES (1, 'ba839f5ed32a42fab544aee16a5f22df');
INSERT INTO `role_webperm_ass` VALUES (1, 'c4ead900038c4fd7bb858a26a5d1e071');
INSERT INTO `role_webperm_ass` VALUES (1, 'e45398f4c8de4257a8671c46681e325e');
COMMIT;

-- ----------------------------
-- Table structure for template
-- ----------------------------
DROP TABLE IF EXISTS `template`;
CREATE TABLE `template` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键,模板ID',
  `tenant_id` varchar(150) NOT NULL COMMENT '租户ID',
  `template_account` varchar(150) NOT NULL COMMENT '模板账号',
  `template_name` varchar(150) NOT NULL COMMENT '模板名称',
  `template_kind` varchar(100) NOT NULL COMMENT '模板类型',
  `content` longtext NOT NULL COMMENT '模板内容',
  `description` text COMMENT '模板介绍',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_tenant_id_template_account` (`tenant_id`,`template_account`),
  CONSTRAINT `fk_template_template_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of template
-- ----------------------------
BEGIN;
INSERT INTO `template` VALUES (7, 'admin', 'temp1', '测试模板1', 'Deployment', '{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"name\":\"nginx-deployment\",\"labels\":{\"app\":\"nginx\"}},\"spec\":{\"replicas\":3,\"selector\":{\"matchLabels\":{\"app\":\"nginx\"}},\"template\":{\"metadata\":{\"labels\":{\"app\":\"nginx\"}},\"spec\":{\"containers\":[{\"name\":\"nginx\",\"image\":\"nginx:1.14.2\",\"ports\":[{\"containerPort\":100}]}]}}}}', '测试模板1的描述', '2020-10-26 07:43:09.947', '2020-10-27 05:49:18.482');
INSERT INTO `template` VALUES (8, 'admin', 'temp2', '测试模板2', 'Deployment', '{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"name\":\"nginx-22222\",\"labels\":{\"app\":\"nginx\"}},\"spec\":{\"replicas\":3,\"selector\":{\"matchLabels\":{\"app\":\"nginx\"}},\"template\":{\"metada\":{\"labels\":{\"app\":\"nginx\"}},\"spec\":{\"containers\":[{\"name\":\"nginx\",\"image\":\"nginx:1.14.2\",\"ports\":[{\"containerPort\":100}]}]}}}}', '测试模板2的描述信息', '2020-10-26 07:43:43.043', '2020-10-30 07:47:02.603');
COMMIT;

-- ----------------------------
-- Table structure for tenant
-- ----------------------------
DROP TABLE IF EXISTS `tenant`;
CREATE TABLE `tenant` (
  `id` varchar(150) NOT NULL COMMENT '主键,租户ID',
  `tenant_name` varchar(150) NOT NULL COMMENT '租户名称',
  `password` varchar(150) DEFAULT NULL COMMENT '密码',
  `salt` varchar(200) DEFAULT NULL COMMENT '密码加密盐',
  `email` varchar(150) DEFAULT '' COMMENT '电子邮箱',
  `description` text COMMENT '用户介绍',
  `last_login_time` datetime(3) DEFAULT NULL COMMENT '最近一次登录时间',
  `last_login_ip` varchar(100) DEFAULT NULL COMMENT '最近一次登录IP',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of tenant
-- ----------------------------
BEGIN;
INSERT INTO `tenant` VALUES ('admin', '主账号', '8d48924c3ac556a83f7411f86f2c50f0a5552cbe31b22d892a06b7767a2ffe8801c1b3ff3c0f5f1d76efd7a673ec1234d27e', 'tv3Eae7B6L', '', '', '2020-11-24 01:47:40.078', '172.16.71.10', '2020-10-21 09:33:22.254', '2020-11-24 01:47:40.152');
INSERT INTO `tenant` VALUES ('builtin_root', '', '9aaf80c9109d2fb7df69bad3e1e4fb59726ffebd992afde001169c82dc3dc3dff1b3c37edfea2d502fa1378a0b3a9e27e84d', '56G4jq673y', '', '', '2020-11-23 06:44:23.985', '172.16.71.17', '2020-10-21 09:33:22.184', '2020-11-23 06:44:24.045');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` varchar(150) NOT NULL COMMENT '主键,用户ID',
  `tenant_id` varchar(150) NOT NULL COMMENT '租户ID',
  `user_account` varchar(150) NOT NULL COMMENT '用户账号',
  `username` varchar(150) NOT NULL COMMENT '用户名称',
  `password` varchar(150) DEFAULT NULL COMMENT '密码',
  `salt` varchar(200) DEFAULT NULL COMMENT '密码加密盐',
  `email` varchar(150) DEFAULT '' COMMENT '电子邮箱',
  `description` text COMMENT '用户介绍',
  `last_login_time` datetime(3) DEFAULT NULL COMMENT '最近一次登录时间',
  `last_login_ip` varchar(100) DEFAULT NULL COMMENT '最近一次登录IP',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `ui_tenant_id_user_account` (`tenant_id`,`user_account`),
  CONSTRAINT `fk_user_user_tenant` FOREIGN KEY (`tenant_id`) REFERENCES `tenant` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES ('FTUSER-3178db1353864cdebe0402cce615ad29', 'admin', 'user2', '用户2', '9ebca807bf3012d63f9969e46717bea0e974cb4aeed92131fcaf60edf7d8df69f08f184c9b907336a4942965c5ee58332965', 'kJU5FJia7x', 'user2@test.com', '用户2的描述信息', NULL, '', '2020-10-21 09:40:47.484', '2020-10-21 09:40:47.484');
INSERT INTO `user` VALUES ('FTUSER-b6b5fd90dd724f70971c86ac8986ea5d', 'admin', 'user3', '用户3', '55b981ee228d2d9ece93c03419a776c3af25980770a1d12510b5f1b679cbd936ba7b502089d222f207b848e8444eb6f3fc35', 'ReiLEVAWe7', '测试修改邮箱@adsf.ccc', '用户3的描述信息', NULL, '', '2020-10-21 09:40:41.040', '2020-10-21 09:50:01.733');
INSERT INTO `user` VALUES ('FTUSER-ec197677ac26430da65cb552642a612b', 'admin', 'user1', '用户1', '4cb2a24355924184c206076663e3dd47889278a1db2c8059e80e6642f044ff24336dc89838b965ff79189fe99984810d3286', 'lUQW1uPkc7', 'user1@1111.com', '用户1的描述信息', '2020-10-23 06:51:27.237', '172.16.71.19', '2020-10-21 09:40:07.274', '2020-10-23 06:51:27.295');
COMMIT;

-- ----------------------------
-- Table structure for user_group_ass
-- ----------------------------
DROP TABLE IF EXISTS `user_group_ass`;
CREATE TABLE `user_group_ass` (
  `group_id` varchar(150) NOT NULL COMMENT '主键,用户组ID',
  `user_id` varchar(150) NOT NULL COMMENT '主键,用户ID',
  PRIMARY KEY (`group_id`,`user_id`),
  KEY `fk_user_group_ass_user` (`user_id`),
  CONSTRAINT `fk_user_group_ass_group` FOREIGN KEY (`group_id`) REFERENCES `group` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `fk_user_group_ass_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of user_group_ass
-- ----------------------------
BEGIN;
INSERT INTO `user_group_ass` VALUES ('FTGROUP-61eedad88efa4fc2a9df83b7111ebd8a', 'FTUSER-3178db1353864cdebe0402cce615ad29');
INSERT INTO `user_group_ass` VALUES ('FTGROUP-61eedad88efa4fc2a9df83b7111ebd8a', 'FTUSER-b6b5fd90dd724f70971c86ac8986ea5d');
INSERT INTO `user_group_ass` VALUES ('FTGROUP-61eedad88efa4fc2a9df83b7111ebd8a', 'FTUSER-ec197677ac26430da65cb552642a612b');
INSERT INTO `user_group_ass` VALUES ('FTGROUP-83f060962723409b9121f43e0555bcff', 'FTUSER-ec197677ac26430da65cb552642a612b');
COMMIT;

-- ----------------------------
-- Table structure for webperm
-- ----------------------------
DROP TABLE IF EXISTS `webperm`;
CREATE TABLE `webperm` (
  `id` varchar(150) NOT NULL COMMENT '主键,前端权限ID',
  `parent_id` varchar(150) DEFAULT NULL COMMENT '前端权限父级ID',
  `name` varchar(191) NOT NULL COMMENT '前端权限名称',
  `path` varchar(191) NOT NULL COMMENT '前端权限路径',
  `resources_sort` bigint NOT NULL COMMENT '排序顺序,1最前,越大越往后',
  `resources_type` varchar(10) NOT NULL COMMENT '权限类型,M:目录,C:资源,F:按钮,H:混合',
  `title` varchar(150) NOT NULL COMMENT '前端权限显示名称',
  `icon` varchar(150) NOT NULL COMMENT '前端权限显示图片',
  `display` bigint NOT NULL COMMENT '是否显示,1:显示,2:隐藏',
  `only_builtin_root` bigint NOT NULL COMMENT '是否只是内置后台管理用户操作,1:是,2:否',
  `create_time` datetime(3) DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime(3) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_webperm_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of webperm
-- ----------------------------
BEGIN;
INSERT INTO `webperm` VALUES ('0125dfeb78c944618e0fe04fb78c1a4f', '8b02c88421c243f6a30ae415593c28bd', 'k8s_timedTaskDetail:tb:detail', 'k8s_load/k8s_timedTaskDetail', 18, 'H', '定时任务详情', '', 1, 2, '2020-11-16 02:03:31.654', '2020-11-16 02:03:31.654');
INSERT INTO `webperm` VALUES ('090c2a954d984d7da92b6fde560472dc', '2ec71c46e79d498d9cf7cc45fc8eebaa', 'k8s_service:tb:detail', 'k8s_service/k8s_serviceDetail', 3, 'H', '服务详情', '', 1, 2, '2020-11-10 09:00:24.860', '2020-11-11 07:34:36.783');
INSERT INTO `webperm` VALUES ('1', '0', 'k8s_cluster', '/k8s_cluster', 1, 'M', '集群', 'colony', 1, 2, '2020-10-21 09:33:23.765', '2020-10-21 09:33:23.765');
INSERT INTO `webperm` VALUES ('1_1', '1', 'k8s_clusterList', 'k8s_clusterList', 1, 'C', '集群列表', 'jiqun', 1, 2, '2020-10-21 09:33:23.781', '2020-10-22 03:29:17.612');
INSERT INTO `webperm` VALUES ('1c61c529be79462da9ec6b98981f0a60', '30d09d0b7ded487c9233445b1b6fbd33', 'k8s_route:tb:detail', 'k8s_route/k8s_routeDetail', 3, 'H', '路由详情', '', 1, 2, '2020-11-10 09:03:28.978', '2020-11-11 07:34:40.335');
INSERT INTO `webperm` VALUES ('2c1bd867a4094d999afca43934fd1084', '1', 'k8s_clusterOverview', 'k8s_clusterOverview', 2, 'C', '集群概览', 'overview', 1, 2, '2020-10-27 09:18:29.102', '2020-10-27 09:18:29.102');
INSERT INTO `webperm` VALUES ('2ec71c46e79d498d9cf7cc45fc8eebaa', '1', 'k8s_service', 'k8s_service', 6, 'C', '服务', 'jichengdefuwu', 1, 2, '2020-11-06 08:59:23.411', '2020-11-06 08:59:23.411');
INSERT INTO `webperm` VALUES ('3', '0', 'k8s_systemSetting', '/k8s_systemSetting', 3, 'M', '系统设置', 'guanli1', 1, 2, '2020-10-21 09:33:23.801', '2020-10-21 09:33:23.801');
INSERT INTO `webperm` VALUES ('3_1', '3', 'k8s_userMg', 'k8s_userMg', 1, 'C', '用户管理', 'yonghu', 1, 2, '2020-10-21 09:33:23.818', '2020-10-22 03:29:54.741');
INSERT INTO `webperm` VALUES ('3_1_1', '3_1', 'k8s_userMg:tb:detail', 'k8s_userMg/k8s_userDetail', 1, 'H', '用户详情', '', 1, 2, '2020-10-21 09:33:23.830', '2020-10-22 03:30:06.547');
INSERT INTO `webperm` VALUES ('3_1_2', '3_1', 'k8s_userMg:add', '#', 2, 'F', '创建', '', 1, 2, '2020-10-21 09:33:23.848', '2020-10-22 03:30:06.548');
INSERT INTO `webperm` VALUES ('3_2', '3', 'k8s_userGroupMg', 'k8s_userGroupMg', 2, 'C', '用户组管理', 'group', 1, 2, '2020-10-21 09:33:23.860', '2020-10-22 03:29:54.743');
INSERT INTO `webperm` VALUES ('3_2_1', '3_2', 'k8s_userGroupMg:tb:detail', 'k8s_userGroupMg/k8s_userGroupDetail', 1, 'H', '用户组详情', '', 1, 2, '2020-10-21 09:33:23.874', '2020-10-22 03:30:10.692');
INSERT INTO `webperm` VALUES ('3_3', '3', 'k8s_roleMg', 'k8s_roleMg', 3, 'C', '角色管理', 'roleAdmin', 1, 2, '2020-10-21 09:33:23.884', '2020-10-22 03:29:54.744');
INSERT INTO `webperm` VALUES ('3_3_1', '3_3', 'k8s_roleMg:tb:detail', 'k8s_roleMg/k8s_roleMgDetail', 1, 'H', '角色管理详情', '', 1, 2, '2020-10-21 09:33:23.894', '2020-10-22 03:29:54.744');
INSERT INTO `webperm` VALUES ('3_4', '3', 'k8s_associateMg', 'k8s_associateMg', 4, 'C', '关联管理', 'guanlian1', 1, 2, '2020-10-21 09:33:23.911', '2020-10-22 03:29:54.745');
INSERT INTO `webperm` VALUES ('3_5', '3', 'k8s_menuMg', 'k8s_menuMg', 5, 'C', '菜单管理', 'yuliushili', 1, 1, '2020-10-21 09:33:23.923', '2020-10-22 03:29:54.745');
INSERT INTO `webperm` VALUES ('30d09d0b7ded487c9233445b1b6fbd33', '1', 'k8s_route', 'k8s_route', 7, 'C', '路由', 'luyoubiao1', 1, 2, '2020-11-06 09:00:09.506', '2020-11-06 09:00:09.506');
INSERT INTO `webperm` VALUES ('3e05f104b88d4531a84d088abc6df56d', '4417f091e42d4ea0bb5749b451e19c1e', 'k8s_storage:yaml:edit', 'k8s_storage/k8s_yamlEdit', 2, 'H', 'Yaml修改', '', 1, 2, '2020-11-18 08:11:45.083', '2020-11-18 08:11:45.083');
INSERT INTO `webperm` VALUES ('4417f091e42d4ea0bb5749b451e19c1e', '1', 'k8s_storage', 'k8s_storage', 8, 'C', '存储', 'cunchu1', 1, 2, '2020-11-13 07:54:31.418', '2020-11-13 07:54:31.418');
INSERT INTO `webperm` VALUES ('49800e4f073f4158981f33be0c7038e9', '8b02c88421c243f6a30ae415593c28bd', 'k8s_taskDetail:tb:detail', 'k8s_load/k8s_taskDetail', 15, 'H', '任务详情', '', 1, 2, '2020-11-16 02:02:20.111', '2020-11-16 02:02:20.111');
INSERT INTO `webperm` VALUES ('4b1c89cc7ad54e53a0f957fe4bf6473a', '2ec71c46e79d498d9cf7cc45fc8eebaa', 'k8s_service:yaml:create', 'k8s_service/k8s_yamlCreate', 1, 'H', 'Yaml创建', '', 1, 2, '2020-11-12 07:10:29.348', '2020-11-12 07:10:29.348');
INSERT INTO `webperm` VALUES ('58f0a536db3949048bdd2405f0fdb029', '1', 'k8s_config', 'k8s_config', 5, 'C', '配置', 'starconfig', 1, 2, '2020-11-06 08:58:33.249', '2020-11-06 08:58:33.249');
INSERT INTO `webperm` VALUES ('614afd34aef04038bd02c2b54f048cbd', '30d09d0b7ded487c9233445b1b6fbd33', 'k8s_route:yaml:create', 'k8s_route/k8s_yamlCreate', 1, 'H', 'Yaml创建', '', 1, 2, '2020-11-13 07:38:47.491', '2020-11-13 07:38:47.491');
INSERT INTO `webperm` VALUES ('6c9c9fce4e9d487ebabad5da63b4aa94', '8b02c88421c243f6a30ae415593c28bd', 'k8s_statelessDetail:tb:detail', 'k8s_load/k8s_statelessDetail', 5, 'H', '无状态详情', '', 1, 2, '2020-10-29 09:25:56.878', '2020-11-11 07:34:28.597');
INSERT INTO `webperm` VALUES ('6ce8d5488991496f871aee0da592e816', '8b02c88421c243f6a30ae415593c28bd', 'k8s_statefulDetail:tb:detail', 'k8s_load/k8s_statefulDetail', 9, 'H', '有状态详情', '', 1, 2, '2020-11-16 01:59:02.219', '2020-11-16 01:59:02.219');
INSERT INTO `webperm` VALUES ('720fb65c7c324f4d8fe724f839128c08', '0', 'k8s_template', '/k8s_template', 2, 'M', '模板', 'moduleMg', 1, 2, '2020-10-23 03:24:44.125', '2020-10-23 03:24:44.125');
INSERT INTO `webperm` VALUES ('787bbc14376c48b1a58f9c73978f68af', '8b02c88421c243f6a30ae415593c28bd', 'k8s_load:yaml:edit', 'k8s_load/k8s_yamlEdit', 2, 'H', 'Yaml修改', '', 1, 2, '2020-10-29 09:37:56.256', '2020-11-11 07:34:28.598');
INSERT INTO `webperm` VALUES ('7fdbb68a604049ceaba4df9a7994e578', '8b02c88421c243f6a30ae415593c28bd', 'k8s_load:normal:edit', 'k8s_load/k8s_normalEdit', 4, 'H', '普通修改', '', 1, 2, '2020-11-09 01:53:28.850', '2020-11-11 07:34:28.598');
INSERT INTO `webperm` VALUES ('8624e370014c46e3bb455a92e68b0c86', '58f0a536db3949048bdd2405f0fdb029', 'k8s_config:yaml:create', 'k8s_config/k8s_yamlCreate', 1, 'H', 'Yaml创建', '', 1, 2, '2020-11-12 02:06:57.738', '2020-11-12 02:06:57.738');
INSERT INTO `webperm` VALUES ('8b02c88421c243f6a30ae415593c28bd', '1', 'k8s_load', 'k8s_load', 4, 'C', '工作负载', 'LVS', 1, 2, '2020-10-23 03:20:29.882', '2020-10-23 03:20:29.882');
INSERT INTO `webperm` VALUES ('97a9799b474847549b7ee60d8440f69b', '720fb65c7c324f4d8fe724f839128c08', 'k8s_templateMg', 'k8s_templateMg', 1, 'C', '模板管理', 'mobanku', 1, 2, '2020-10-23 03:25:56.726', '2020-10-23 05:47:14.935');
INSERT INTO `webperm` VALUES ('9aedcb6cbf8442afa8206542c535e1dc', '58f0a536db3949048bdd2405f0fdb029', 'k8s_dictDetail:tb:detail', 'k8s_config/k8s_dictDetail', 4, 'H', '保密字典详情', '', 1, 2, '2020-11-10 07:05:25.326', '2020-11-11 07:34:32.968');
INSERT INTO `webperm` VALUES ('9ea21d65985c4b2ea9f856e47fc51c92', '58f0a536db3949048bdd2405f0fdb029', 'k8s_configDetail:tb:detail', 'k8s_config/k8s_configDetail', 3, 'H', '配置项详情', '', 1, 2, '2020-11-10 07:02:51.937', '2020-11-11 07:34:32.969');
INSERT INTO `webperm` VALUES ('a18d07bb29c84ad5befb20353f8dfa08', '8b02c88421c243f6a30ae415593c28bd', 'k8s_containerDetail:tb:detail', 'k8s_load/k8s_containerDetail', 6, 'H', '容器组详情', '', 1, 2, '2020-11-06 02:57:20.239', '2020-11-11 07:34:28.599');
INSERT INTO `webperm` VALUES ('aec53260e7104746a4aac26fc36b913c', '30d09d0b7ded487c9233445b1b6fbd33', 'k8s_route:yaml:edit', 'k8s_route/k8s_yamlEdit', 2, 'H', 'Yaml修改', '', 1, 2, '2020-11-13 07:40:05.526', '2020-11-13 07:40:05.526');
INSERT INTO `webperm` VALUES ('ba839f5ed32a42fab544aee16a5f22df', '8b02c88421c243f6a30ae415593c28bd', 'k8s_load:yaml:create', 'k8s_load/k8s_yamlCreate', 1, 'H', 'Yaml创建', '', 1, 2, '2020-10-29 09:36:59.823', '2020-11-11 07:34:28.599');
INSERT INTO `webperm` VALUES ('c4ead900038c4fd7bb858a26a5d1e071', '97a9799b474847549b7ee60d8440f69b', 'k8s_templateDetail:tb:detail', 'k8s_templateMg/k8s_templateDetail', 1, 'H', '查看模板', '', 1, 2, '2020-10-26 07:48:07.696', '2020-10-26 07:48:07.696');
INSERT INTO `webperm` VALUES ('d2efb62501c04a1888dd8f92971e2070', '8b02c88421c243f6a30ae415593c28bd', 'k8s_load:normal:create', 'k8s_load/k8s_normalCreate', 3, 'H', '普通创建', '', 1, 2, '2020-11-09 01:52:06.534', '2020-11-11 07:34:28.600');
INSERT INTO `webperm` VALUES ('dda01b0a743744828aa30f9ed9a830dd', '4417f091e42d4ea0bb5749b451e19c1e', 'k8s_storage:yaml:create', 'k8s_storage/k8s_yamlCreate', 1, 'H', 'Yaml创建', '', 1, 2, '2020-11-18 08:10:55.121', '2020-11-18 08:10:55.121');
INSERT INTO `webperm` VALUES ('e45398f4c8de4257a8671c46681e325e', '1', 'k8s_namespace', 'k8s_namespace', 3, 'C', '命名空间', 'card', 1, 2, '2020-10-23 03:19:37.426', '2020-10-23 03:19:37.426');
INSERT INTO `webperm` VALUES ('f030361ea8264cae807680e25f285023', '8b02c88421c243f6a30ae415593c28bd', 'k8s_processDetail:tb:detail', 'k8s_load/k8s_processDetail', 12, 'H', '进程守护集详情', '', 1, 2, '2020-11-16 02:00:49.243', '2020-11-16 02:00:49.243');
INSERT INTO `webperm` VALUES ('f9885513c6fb4dd7aa51b76b7e0a2be2', '58f0a536db3949048bdd2405f0fdb029', 'k8s_config:yaml:edit', 'k8s_config/k8s_yamlEdit', 2, 'H', 'Yaml修改', '', 1, 2, '2020-11-12 02:07:58.889', '2020-11-12 02:07:58.889');
INSERT INTO `webperm` VALUES ('fa4fdaab33e14722a7a20c9dfe509f4a', '2ec71c46e79d498d9cf7cc45fc8eebaa', 'k8s_service:yaml:edit', 'k8s_service/k8s_yamlEdit', 2, 'H', 'Yaml修改', '', 1, 2, '2020-11-12 07:11:09.971', '2020-11-12 07:11:09.971');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
