// source: heyrenee/v1/messages/patient.proto
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

var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
goog.object.extend(proto, google_protobuf_timestamp_pb);
var heyrenee_v1_options_auth_pb = require('../../../heyrenee/v1/options/auth_pb.js');
goog.object.extend(proto, heyrenee_v1_options_auth_pb);
var heyrenee_v1_enums_language_pb = require('../../../heyrenee/v1/enums/language_pb.js');
goog.object.extend(proto, heyrenee_v1_enums_language_pb);
goog.exportSymbol('proto.heyrenee.v1.messages.Ethnicity', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.MaritalStatus', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.Patient', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.Sex', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.heyrenee.v1.messages.Patient = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.heyrenee.v1.messages.Patient.repeatedFields_, null);
};
goog.inherits(proto.heyrenee.v1.messages.Patient, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.messages.Patient.displayName = 'proto.heyrenee.v1.messages.Patient';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.heyrenee.v1.messages.Patient.repeatedFields_ = [9];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.heyrenee.v1.messages.Patient.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.messages.Patient.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.messages.Patient} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.Patient.toObject = function(includeInstance, msg) {
  var f, obj = {
    patientId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    firstName: jspb.Message.getFieldWithDefault(msg, 2, ""),
    middleName: jspb.Message.getFieldWithDefault(msg, 3, ""),
    lastName: jspb.Message.getFieldWithDefault(msg, 4, ""),
    primaryPhone: jspb.Message.getFieldWithDefault(msg, 5, ""),
    dateOfBirth: (f = msg.getDateOfBirth()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f),
    sex: jspb.Message.getFieldWithDefault(msg, 7, 0),
    email: jspb.Message.getFieldWithDefault(msg, 8, ""),
    addressLinesList: (f = jspb.Message.getRepeatedField(msg, 9)) == null ? undefined : f,
    city: jspb.Message.getFieldWithDefault(msg, 10, ""),
    state: jspb.Message.getFieldWithDefault(msg, 11, ""),
    zip: jspb.Message.getFieldWithDefault(msg, 12, ""),
    preferredName: jspb.Message.getFieldWithDefault(msg, 13, ""),
    maritalStatus: jspb.Message.getFieldWithDefault(msg, 14, 0),
    preferredLanguage: jspb.Message.getFieldWithDefault(msg, 15, 0),
    ethnicity: jspb.Message.getFieldWithDefault(msg, 16, 0),
    mobilePhone: jspb.Message.getFieldWithDefault(msg, 17, ""),
    otherPhone: jspb.Message.getFieldWithDefault(msg, 18, ""),
    notes: jspb.Message.getFieldWithDefault(msg, 19, ""),
    namePronunciation: jspb.Message.getFieldWithDefault(msg, 20, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.heyrenee.v1.messages.Patient}
 */
proto.heyrenee.v1.messages.Patient.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.messages.Patient;
  return proto.heyrenee.v1.messages.Patient.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.messages.Patient} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.messages.Patient}
 */
proto.heyrenee.v1.messages.Patient.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setPatientId(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setFirstName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setMiddleName(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setLastName(value);
      break;
    case 5:
      var value = /** @type {string} */ (reader.readString());
      msg.setPrimaryPhone(value);
      break;
    case 6:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setDateOfBirth(value);
      break;
    case 7:
      var value = /** @type {!proto.heyrenee.v1.messages.Sex} */ (reader.readEnum());
      msg.setSex(value);
      break;
    case 8:
      var value = /** @type {string} */ (reader.readString());
      msg.setEmail(value);
      break;
    case 9:
      var value = /** @type {string} */ (reader.readString());
      msg.addAddressLines(value);
      break;
    case 10:
      var value = /** @type {string} */ (reader.readString());
      msg.setCity(value);
      break;
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setState(value);
      break;
    case 12:
      var value = /** @type {string} */ (reader.readString());
      msg.setZip(value);
      break;
    case 13:
      var value = /** @type {string} */ (reader.readString());
      msg.setPreferredName(value);
      break;
    case 14:
      var value = /** @type {!proto.heyrenee.v1.messages.MaritalStatus} */ (reader.readEnum());
      msg.setMaritalStatus(value);
      break;
    case 15:
      var value = /** @type {!proto.heyrenee.v1.enums.Language} */ (reader.readEnum());
      msg.setPreferredLanguage(value);
      break;
    case 16:
      var value = /** @type {!proto.heyrenee.v1.messages.Ethnicity} */ (reader.readEnum());
      msg.setEthnicity(value);
      break;
    case 17:
      var value = /** @type {string} */ (reader.readString());
      msg.setMobilePhone(value);
      break;
    case 18:
      var value = /** @type {string} */ (reader.readString());
      msg.setOtherPhone(value);
      break;
    case 19:
      var value = /** @type {string} */ (reader.readString());
      msg.setNotes(value);
      break;
    case 20:
      var value = /** @type {string} */ (reader.readString());
      msg.setNamePronunciation(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.heyrenee.v1.messages.Patient.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.messages.Patient.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.messages.Patient} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.Patient.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPatientId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getFirstName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getMiddleName();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getLastName();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getPrimaryPhone();
  if (f.length > 0) {
    writer.writeString(
      5,
      f
    );
  }
  f = message.getDateOfBirth();
  if (f != null) {
    writer.writeMessage(
      6,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
  f = message.getSex();
  if (f !== 0.0) {
    writer.writeEnum(
      7,
      f
    );
  }
  f = message.getEmail();
  if (f.length > 0) {
    writer.writeString(
      8,
      f
    );
  }
  f = message.getAddressLinesList();
  if (f.length > 0) {
    writer.writeRepeatedString(
      9,
      f
    );
  }
  f = message.getCity();
  if (f.length > 0) {
    writer.writeString(
      10,
      f
    );
  }
  f = message.getState();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getZip();
  if (f.length > 0) {
    writer.writeString(
      12,
      f
    );
  }
  f = message.getPreferredName();
  if (f.length > 0) {
    writer.writeString(
      13,
      f
    );
  }
  f = message.getMaritalStatus();
  if (f !== 0.0) {
    writer.writeEnum(
      14,
      f
    );
  }
  f = message.getPreferredLanguage();
  if (f !== 0.0) {
    writer.writeEnum(
      15,
      f
    );
  }
  f = message.getEthnicity();
  if (f !== 0.0) {
    writer.writeEnum(
      16,
      f
    );
  }
  f = message.getMobilePhone();
  if (f.length > 0) {
    writer.writeString(
      17,
      f
    );
  }
  f = message.getOtherPhone();
  if (f.length > 0) {
    writer.writeString(
      18,
      f
    );
  }
  f = message.getNotes();
  if (f.length > 0) {
    writer.writeString(
      19,
      f
    );
  }
  f = message.getNamePronunciation();
  if (f.length > 0) {
    writer.writeString(
      20,
      f
    );
  }
};


/**
 * optional string patient_id = 1;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getPatientId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setPatientId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string first_name = 2;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getFirstName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setFirstName = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string middle_name = 3;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getMiddleName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setMiddleName = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string last_name = 4;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getLastName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setLastName = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional string primary_phone = 5;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getPrimaryPhone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 5, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setPrimaryPhone = function(value) {
  return jspb.Message.setProto3StringField(this, 5, value);
};


/**
 * optional google.protobuf.Timestamp date_of_birth = 6;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.heyrenee.v1.messages.Patient.prototype.getDateOfBirth = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 6));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
*/
proto.heyrenee.v1.messages.Patient.prototype.setDateOfBirth = function(value) {
  return jspb.Message.setWrapperField(this, 6, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.clearDateOfBirth = function() {
  return this.setDateOfBirth(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.Patient.prototype.hasDateOfBirth = function() {
  return jspb.Message.getField(this, 6) != null;
};


/**
 * optional Sex sex = 7;
 * @return {!proto.heyrenee.v1.messages.Sex}
 */
proto.heyrenee.v1.messages.Patient.prototype.getSex = function() {
  return /** @type {!proto.heyrenee.v1.messages.Sex} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.Sex} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setSex = function(value) {
  return jspb.Message.setProto3EnumField(this, 7, value);
};


/**
 * optional string email = 8;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getEmail = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 8, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setEmail = function(value) {
  return jspb.Message.setProto3StringField(this, 8, value);
};


/**
 * repeated string address_lines = 9;
 * @return {!Array<string>}
 */
proto.heyrenee.v1.messages.Patient.prototype.getAddressLinesList = function() {
  return /** @type {!Array<string>} */ (jspb.Message.getRepeatedField(this, 9));
};


/**
 * @param {!Array<string>} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setAddressLinesList = function(value) {
  return jspb.Message.setField(this, 9, value || []);
};


/**
 * @param {string} value
 * @param {number=} opt_index
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.addAddressLines = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 9, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.clearAddressLinesList = function() {
  return this.setAddressLinesList([]);
};


/**
 * optional string city = 10;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getCity = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 10, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setCity = function(value) {
  return jspb.Message.setProto3StringField(this, 10, value);
};


/**
 * optional string state = 11;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getState = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setState = function(value) {
  return jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * optional string zip = 12;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getZip = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 12, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setZip = function(value) {
  return jspb.Message.setProto3StringField(this, 12, value);
};


/**
 * optional string preferred_name = 13;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getPreferredName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 13, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setPreferredName = function(value) {
  return jspb.Message.setProto3StringField(this, 13, value);
};


/**
 * optional MaritalStatus marital_status = 14;
 * @return {!proto.heyrenee.v1.messages.MaritalStatus}
 */
proto.heyrenee.v1.messages.Patient.prototype.getMaritalStatus = function() {
  return /** @type {!proto.heyrenee.v1.messages.MaritalStatus} */ (jspb.Message.getFieldWithDefault(this, 14, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.MaritalStatus} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setMaritalStatus = function(value) {
  return jspb.Message.setProto3EnumField(this, 14, value);
};


/**
 * optional heyrenee.v1.enums.Language preferred_language = 15;
 * @return {!proto.heyrenee.v1.enums.Language}
 */
proto.heyrenee.v1.messages.Patient.prototype.getPreferredLanguage = function() {
  return /** @type {!proto.heyrenee.v1.enums.Language} */ (jspb.Message.getFieldWithDefault(this, 15, 0));
};


/**
 * @param {!proto.heyrenee.v1.enums.Language} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setPreferredLanguage = function(value) {
  return jspb.Message.setProto3EnumField(this, 15, value);
};


/**
 * optional Ethnicity ethnicity = 16;
 * @return {!proto.heyrenee.v1.messages.Ethnicity}
 */
proto.heyrenee.v1.messages.Patient.prototype.getEthnicity = function() {
  return /** @type {!proto.heyrenee.v1.messages.Ethnicity} */ (jspb.Message.getFieldWithDefault(this, 16, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.Ethnicity} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setEthnicity = function(value) {
  return jspb.Message.setProto3EnumField(this, 16, value);
};


/**
 * optional string mobile_phone = 17;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getMobilePhone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 17, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setMobilePhone = function(value) {
  return jspb.Message.setProto3StringField(this, 17, value);
};


/**
 * optional string other_phone = 18;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getOtherPhone = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 18, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setOtherPhone = function(value) {
  return jspb.Message.setProto3StringField(this, 18, value);
};


/**
 * optional string notes = 19;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getNotes = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 19, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setNotes = function(value) {
  return jspb.Message.setProto3StringField(this, 19, value);
};


/**
 * optional string name_pronunciation = 20;
 * @return {string}
 */
proto.heyrenee.v1.messages.Patient.prototype.getNamePronunciation = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 20, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.Patient} returns this
 */
proto.heyrenee.v1.messages.Patient.prototype.setNamePronunciation = function(value) {
  return jspb.Message.setProto3StringField(this, 20, value);
};


/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.Sex = {
  SEX_UNSPECIFIED: 0,
  SEX_MALE: 1,
  SEX_FEMALE: 2
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.MaritalStatus = {
  MARITAL_STATUS_UNSPECIFIED: 0,
  MARITAL_STATUS_SINGLE: 1,
  MARITAL_STATUS_MARRIED: 2
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.Ethnicity = {
  ETHNICITY_UNSPECIFIED: 0,
  ETHNICITY_ASIAN_AMERICAN: 1,
  ETHNICITY_BLACK_OR_AFRICAN_AMERICAN: 2,
  ETHNICITY_WHITE_OR_EUROPEAN_AMERICAN: 3,
  ETHNICITY_AMERICAN_INDIAN_OR_ALASKA_NATIVE: 4,
  ETHNICITY_NATIVE_HAWAIIAN_OR_PACIFIC_ISLANDER: 5
};

goog.object.extend(exports, proto.heyrenee.v1.messages);