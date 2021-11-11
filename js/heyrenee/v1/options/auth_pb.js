// source: heyrenee/v1/options/auth.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {missingRequire} reports error on implicit type usages.
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var google_protobuf_descriptor_pb = require('google-protobuf/google/protobuf/descriptor_pb.js');
goog.object.extend(proto, google_protobuf_descriptor_pb);
goog.exportSymbol('proto.heyrenee.v1.options.FieldAuthorization', null, global);
goog.exportSymbol('proto.heyrenee.v1.options.fieldAuthorization', null, global);
/**
 * @enum {number}
 */
proto.heyrenee.v1.options.FieldAuthorization = {
  FIELD_AUTHORIZATION_NONE: 0,
  FIELD_AUTHORIZATION_PATIENT: 1,
  FIELD_AUTHORIZATION_CAREGIVER: 2,
  FIELD_AUTHORIZATION_CONCIERGE: 3,
  FIELD_AUTHORIZATION_PRESCRIPTION: 4,
  FIELD_AUTHORIZATION_RPM_SCHEDULE: 5
};


/**
 * A tuple of {field number, class constructor} for the extension
 * field named `fieldAuthorization`.
 * @type {!jspb.ExtensionFieldInfo<!proto.heyrenee.v1.options.FieldAuthorization>}
 */
proto.heyrenee.v1.options.fieldAuthorization = new jspb.ExtensionFieldInfo(
    50000,
    {fieldAuthorization: 0},
    null,
     /** @type {?function((boolean|undefined),!jspb.Message=): !Object} */ (
         null),
    0);

google_protobuf_descriptor_pb.FieldOptions.extensionsBinary[50000] = new jspb.ExtensionFieldBinaryInfo(
    proto.heyrenee.v1.options.fieldAuthorization,
    jspb.BinaryReader.prototype.readEnum,
    jspb.BinaryWriter.prototype.writeEnum,
    undefined,
    undefined,
    false);
// This registers the extension field with the extended class, so that
// toObject() will function correctly.
google_protobuf_descriptor_pb.FieldOptions.extensions[50000] = proto.heyrenee.v1.options.fieldAuthorization;

goog.object.extend(exports, proto.heyrenee.v1.options);
