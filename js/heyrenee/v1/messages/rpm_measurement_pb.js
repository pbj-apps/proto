// source: heyrenee/v1/messages/rpm_measurement.proto
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
goog.exportSymbol('proto.heyrenee.v1.messages.RpmMeasurement', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.RpmMeasurementResult', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.RpmMeasurementResultType', null, global);
goog.exportSymbol('proto.heyrenee.v1.messages.RpmMeasurementResultUnit', null, global);
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
proto.heyrenee.v1.messages.RpmMeasurement = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.heyrenee.v1.messages.RpmMeasurement.repeatedFields_, null);
};
goog.inherits(proto.heyrenee.v1.messages.RpmMeasurement, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.messages.RpmMeasurement.displayName = 'proto.heyrenee.v1.messages.RpmMeasurement';
}
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
proto.heyrenee.v1.messages.RpmMeasurementResult = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.heyrenee.v1.messages.RpmMeasurementResult, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.messages.RpmMeasurementResult.displayName = 'proto.heyrenee.v1.messages.RpmMeasurementResult';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.heyrenee.v1.messages.RpmMeasurement.repeatedFields_ = [4];



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
proto.heyrenee.v1.messages.RpmMeasurement.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.messages.RpmMeasurement.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.messages.RpmMeasurement} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.RpmMeasurement.toObject = function(includeInstance, msg) {
  var f, obj = {
    patientId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    rpmScheduleId: jspb.Message.getFieldWithDefault(msg, 2, ""),
    rpmMeasurementId: jspb.Message.getFieldWithDefault(msg, 3, ""),
    rpmMeasurementResultsList: jspb.Message.toObjectList(msg.getRpmMeasurementResultsList(),
    proto.heyrenee.v1.messages.RpmMeasurementResult.toObject, includeInstance),
    timeTaken: (f = msg.getTimeTaken()) && google_protobuf_timestamp_pb.Timestamp.toObject(includeInstance, f)
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
 * @return {!proto.heyrenee.v1.messages.RpmMeasurement}
 */
proto.heyrenee.v1.messages.RpmMeasurement.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.messages.RpmMeasurement;
  return proto.heyrenee.v1.messages.RpmMeasurement.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.messages.RpmMeasurement} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.messages.RpmMeasurement}
 */
proto.heyrenee.v1.messages.RpmMeasurement.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setRpmScheduleId(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setRpmMeasurementId(value);
      break;
    case 4:
      var value = new proto.heyrenee.v1.messages.RpmMeasurementResult;
      reader.readMessage(value,proto.heyrenee.v1.messages.RpmMeasurementResult.deserializeBinaryFromReader);
      msg.addRpmMeasurementResults(value);
      break;
    case 5:
      var value = new google_protobuf_timestamp_pb.Timestamp;
      reader.readMessage(value,google_protobuf_timestamp_pb.Timestamp.deserializeBinaryFromReader);
      msg.setTimeTaken(value);
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
proto.heyrenee.v1.messages.RpmMeasurement.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.messages.RpmMeasurement.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.messages.RpmMeasurement} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.RpmMeasurement.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPatientId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getRpmScheduleId();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getRpmMeasurementId();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getRpmMeasurementResultsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      4,
      f,
      proto.heyrenee.v1.messages.RpmMeasurementResult.serializeBinaryToWriter
    );
  }
  f = message.getTimeTaken();
  if (f != null) {
    writer.writeMessage(
      5,
      f,
      google_protobuf_timestamp_pb.Timestamp.serializeBinaryToWriter
    );
  }
};


/**
 * optional string patient_id = 1;
 * @return {string}
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.getPatientId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.RpmMeasurement} returns this
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.setPatientId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string rpm_schedule_id = 2;
 * @return {string}
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.getRpmScheduleId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.RpmMeasurement} returns this
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.setRpmScheduleId = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string rpm_measurement_id = 3;
 * @return {string}
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.getRpmMeasurementId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.RpmMeasurement} returns this
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.setRpmMeasurementId = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * repeated RpmMeasurementResult rpm_measurement_results = 4;
 * @return {!Array<!proto.heyrenee.v1.messages.RpmMeasurementResult>}
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.getRpmMeasurementResultsList = function() {
  return /** @type{!Array<!proto.heyrenee.v1.messages.RpmMeasurementResult>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.heyrenee.v1.messages.RpmMeasurementResult, 4));
};


/**
 * @param {!Array<!proto.heyrenee.v1.messages.RpmMeasurementResult>} value
 * @return {!proto.heyrenee.v1.messages.RpmMeasurement} returns this
*/
proto.heyrenee.v1.messages.RpmMeasurement.prototype.setRpmMeasurementResultsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 4, value);
};


/**
 * @param {!proto.heyrenee.v1.messages.RpmMeasurementResult=} opt_value
 * @param {number=} opt_index
 * @return {!proto.heyrenee.v1.messages.RpmMeasurementResult}
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.addRpmMeasurementResults = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 4, opt_value, proto.heyrenee.v1.messages.RpmMeasurementResult, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.heyrenee.v1.messages.RpmMeasurement} returns this
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.clearRpmMeasurementResultsList = function() {
  return this.setRpmMeasurementResultsList([]);
};


/**
 * optional google.protobuf.Timestamp time_taken = 5;
 * @return {?proto.google.protobuf.Timestamp}
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.getTimeTaken = function() {
  return /** @type{?proto.google.protobuf.Timestamp} */ (
    jspb.Message.getWrapperField(this, google_protobuf_timestamp_pb.Timestamp, 5));
};


/**
 * @param {?proto.google.protobuf.Timestamp|undefined} value
 * @return {!proto.heyrenee.v1.messages.RpmMeasurement} returns this
*/
proto.heyrenee.v1.messages.RpmMeasurement.prototype.setTimeTaken = function(value) {
  return jspb.Message.setWrapperField(this, 5, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.messages.RpmMeasurement} returns this
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.clearTimeTaken = function() {
  return this.setTimeTaken(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.messages.RpmMeasurement.prototype.hasTimeTaken = function() {
  return jspb.Message.getField(this, 5) != null;
};





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
proto.heyrenee.v1.messages.RpmMeasurementResult.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.messages.RpmMeasurementResult.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.messages.RpmMeasurementResult} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.toObject = function(includeInstance, msg) {
  var f, obj = {
    rpmMeasurementResultId: jspb.Message.getFieldWithDefault(msg, 1, ""),
    value: jspb.Message.getFloatingPointFieldWithDefault(msg, 2, 0.0),
    rpmMeasurementResultUnit: jspb.Message.getFieldWithDefault(msg, 3, 0),
    rpmMeasurementResultType: jspb.Message.getFieldWithDefault(msg, 4, 0)
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
 * @return {!proto.heyrenee.v1.messages.RpmMeasurementResult}
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.messages.RpmMeasurementResult;
  return proto.heyrenee.v1.messages.RpmMeasurementResult.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.messages.RpmMeasurementResult} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.messages.RpmMeasurementResult}
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setRpmMeasurementResultId(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setValue(value);
      break;
    case 3:
      var value = /** @type {!proto.heyrenee.v1.messages.RpmMeasurementResultUnit} */ (reader.readEnum());
      msg.setRpmMeasurementResultUnit(value);
      break;
    case 4:
      var value = /** @type {!proto.heyrenee.v1.messages.RpmMeasurementResultType} */ (reader.readEnum());
      msg.setRpmMeasurementResultType(value);
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
proto.heyrenee.v1.messages.RpmMeasurementResult.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.messages.RpmMeasurementResult.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.messages.RpmMeasurementResult} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getRpmMeasurementResultId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getValue();
  if (f !== 0.0) {
    writer.writeDouble(
      2,
      f
    );
  }
  f = message.getRpmMeasurementResultUnit();
  if (f !== 0.0) {
    writer.writeEnum(
      3,
      f
    );
  }
  f = message.getRpmMeasurementResultType();
  if (f !== 0.0) {
    writer.writeEnum(
      4,
      f
    );
  }
};


/**
 * optional string rpm_measurement_result_id = 1;
 * @return {string}
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.prototype.getRpmMeasurementResultId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.messages.RpmMeasurementResult} returns this
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.prototype.setRpmMeasurementResultId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional double value = 2;
 * @return {number}
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.prototype.getValue = function() {
  return /** @type {number} */ (jspb.Message.getFloatingPointFieldWithDefault(this, 2, 0.0));
};


/**
 * @param {number} value
 * @return {!proto.heyrenee.v1.messages.RpmMeasurementResult} returns this
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.prototype.setValue = function(value) {
  return jspb.Message.setProto3FloatField(this, 2, value);
};


/**
 * optional RpmMeasurementResultUnit rpm_measurement_result_unit = 3;
 * @return {!proto.heyrenee.v1.messages.RpmMeasurementResultUnit}
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.prototype.getRpmMeasurementResultUnit = function() {
  return /** @type {!proto.heyrenee.v1.messages.RpmMeasurementResultUnit} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.RpmMeasurementResultUnit} value
 * @return {!proto.heyrenee.v1.messages.RpmMeasurementResult} returns this
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.prototype.setRpmMeasurementResultUnit = function(value) {
  return jspb.Message.setProto3EnumField(this, 3, value);
};


/**
 * optional RpmMeasurementResultType rpm_measurement_result_type = 4;
 * @return {!proto.heyrenee.v1.messages.RpmMeasurementResultType}
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.prototype.getRpmMeasurementResultType = function() {
  return /** @type {!proto.heyrenee.v1.messages.RpmMeasurementResultType} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/**
 * @param {!proto.heyrenee.v1.messages.RpmMeasurementResultType} value
 * @return {!proto.heyrenee.v1.messages.RpmMeasurementResult} returns this
 */
proto.heyrenee.v1.messages.RpmMeasurementResult.prototype.setRpmMeasurementResultType = function(value) {
  return jspb.Message.setProto3EnumField(this, 4, value);
};


/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.RpmMeasurementResultUnit = {
  RPM_MEASUREMENT_RESULT_UNIT_UNSPECIFIED: 0,
  RPM_MEASUREMENT_RESULT_UNIT_BPM: 1,
  RPM_MEASUREMENT_RESULT_UNIT_LBS: 2,
  RPM_MEASUREMENT_RESULT_UNIT_PERCENTAGE: 3,
  RPM_MEASUREMENT_RESULT_UNIT_MG_DL: 4,
  RPM_MEASUREMENT_RESULT_UNIT_MM_HG: 5,
  RPM_MEASUREMENT_RESULT_UNIT_CELSIUS: 6
};

/**
 * @enum {number}
 */
proto.heyrenee.v1.messages.RpmMeasurementResultType = {
  RPM_MEASUREMENT_RESULT_TYPE_UNSPECIFIED: 0,
  RPM_MEASUREMENT_RESULT_TYPE_WEIGHT: 1,
  RPM_MEASUREMENT_RESULT_TYPE_SPO2: 2,
  RPM_MEASUREMENT_RESULT_TYPE_BPM: 3,
  RPM_MEASUREMENT_RESULT_TYPE_GLUCOSE: 4,
  RPM_MEASUREMENT_RESULT_TYPE_DIA: 5,
  RPM_MEASUREMENT_RESULT_TYPE_SYS: 6,
  RPM_MEASUREMENT_RESULT_TYPE_TEMP: 7
};

goog.object.extend(exports, proto.heyrenee.v1.messages);
