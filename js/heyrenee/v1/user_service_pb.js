// source: heyrenee/v1/user_service.proto
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

var heyrenee_v1_options_auth_pb = require('../../heyrenee/v1/options/auth_pb.js');
goog.object.extend(proto, heyrenee_v1_options_auth_pb);
var heyrenee_v1_messages_patient_pb = require('../../heyrenee/v1/messages/patient_pb.js');
goog.object.extend(proto, heyrenee_v1_messages_patient_pb);
var heyrenee_v1_messages_caregiver_pb = require('../../heyrenee/v1/messages/caregiver_pb.js');
goog.object.extend(proto, heyrenee_v1_messages_caregiver_pb);
goog.exportSymbol('proto.heyrenee.v1.CreateCaregiverRequest', null, global);
goog.exportSymbol('proto.heyrenee.v1.CreatePatientRequest', null, global);
goog.exportSymbol('proto.heyrenee.v1.GetCaregiverRequest', null, global);
goog.exportSymbol('proto.heyrenee.v1.GetPatientRequest', null, global);
goog.exportSymbol('proto.heyrenee.v1.UpdateCaregiverRequest', null, global);
goog.exportSymbol('proto.heyrenee.v1.UpdatePatientRequest', null, global);
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
proto.heyrenee.v1.CreatePatientRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.heyrenee.v1.CreatePatientRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.CreatePatientRequest.displayName = 'proto.heyrenee.v1.CreatePatientRequest';
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
proto.heyrenee.v1.GetPatientRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.heyrenee.v1.GetPatientRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.GetPatientRequest.displayName = 'proto.heyrenee.v1.GetPatientRequest';
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
proto.heyrenee.v1.UpdatePatientRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.heyrenee.v1.UpdatePatientRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.UpdatePatientRequest.displayName = 'proto.heyrenee.v1.UpdatePatientRequest';
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
proto.heyrenee.v1.CreateCaregiverRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.heyrenee.v1.CreateCaregiverRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.CreateCaregiverRequest.displayName = 'proto.heyrenee.v1.CreateCaregiverRequest';
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
proto.heyrenee.v1.GetCaregiverRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.heyrenee.v1.GetCaregiverRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.GetCaregiverRequest.displayName = 'proto.heyrenee.v1.GetCaregiverRequest';
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
proto.heyrenee.v1.UpdateCaregiverRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.heyrenee.v1.UpdateCaregiverRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.heyrenee.v1.UpdateCaregiverRequest.displayName = 'proto.heyrenee.v1.UpdateCaregiverRequest';
}



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
proto.heyrenee.v1.CreatePatientRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.CreatePatientRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.CreatePatientRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.CreatePatientRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    patient: (f = msg.getPatient()) && heyrenee_v1_messages_patient_pb.Patient.toObject(includeInstance, f),
    password: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.heyrenee.v1.CreatePatientRequest}
 */
proto.heyrenee.v1.CreatePatientRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.CreatePatientRequest;
  return proto.heyrenee.v1.CreatePatientRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.CreatePatientRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.CreatePatientRequest}
 */
proto.heyrenee.v1.CreatePatientRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new heyrenee_v1_messages_patient_pb.Patient;
      reader.readMessage(value,heyrenee_v1_messages_patient_pb.Patient.deserializeBinaryFromReader);
      msg.setPatient(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
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
proto.heyrenee.v1.CreatePatientRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.CreatePatientRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.CreatePatientRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.CreatePatientRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPatient();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      heyrenee_v1_messages_patient_pb.Patient.serializeBinaryToWriter
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional messages.Patient patient = 1;
 * @return {?proto.heyrenee.v1.messages.Patient}
 */
proto.heyrenee.v1.CreatePatientRequest.prototype.getPatient = function() {
  return /** @type{?proto.heyrenee.v1.messages.Patient} */ (
    jspb.Message.getWrapperField(this, heyrenee_v1_messages_patient_pb.Patient, 1));
};


/**
 * @param {?proto.heyrenee.v1.messages.Patient|undefined} value
 * @return {!proto.heyrenee.v1.CreatePatientRequest} returns this
*/
proto.heyrenee.v1.CreatePatientRequest.prototype.setPatient = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.CreatePatientRequest} returns this
 */
proto.heyrenee.v1.CreatePatientRequest.prototype.clearPatient = function() {
  return this.setPatient(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.CreatePatientRequest.prototype.hasPatient = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string password = 2;
 * @return {string}
 */
proto.heyrenee.v1.CreatePatientRequest.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.CreatePatientRequest} returns this
 */
proto.heyrenee.v1.CreatePatientRequest.prototype.setPassword = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
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
proto.heyrenee.v1.GetPatientRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.GetPatientRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.GetPatientRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.GetPatientRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    patientId: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.heyrenee.v1.GetPatientRequest}
 */
proto.heyrenee.v1.GetPatientRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.GetPatientRequest;
  return proto.heyrenee.v1.GetPatientRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.GetPatientRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.GetPatientRequest}
 */
proto.heyrenee.v1.GetPatientRequest.deserializeBinaryFromReader = function(msg, reader) {
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
proto.heyrenee.v1.GetPatientRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.GetPatientRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.GetPatientRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.GetPatientRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPatientId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string patient_id = 1;
 * @return {string}
 */
proto.heyrenee.v1.GetPatientRequest.prototype.getPatientId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.GetPatientRequest} returns this
 */
proto.heyrenee.v1.GetPatientRequest.prototype.setPatientId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
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
proto.heyrenee.v1.UpdatePatientRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.UpdatePatientRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.UpdatePatientRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.UpdatePatientRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    patient: (f = msg.getPatient()) && heyrenee_v1_messages_patient_pb.Patient.toObject(includeInstance, f)
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
 * @return {!proto.heyrenee.v1.UpdatePatientRequest}
 */
proto.heyrenee.v1.UpdatePatientRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.UpdatePatientRequest;
  return proto.heyrenee.v1.UpdatePatientRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.UpdatePatientRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.UpdatePatientRequest}
 */
proto.heyrenee.v1.UpdatePatientRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new heyrenee_v1_messages_patient_pb.Patient;
      reader.readMessage(value,heyrenee_v1_messages_patient_pb.Patient.deserializeBinaryFromReader);
      msg.setPatient(value);
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
proto.heyrenee.v1.UpdatePatientRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.UpdatePatientRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.UpdatePatientRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.UpdatePatientRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPatient();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      heyrenee_v1_messages_patient_pb.Patient.serializeBinaryToWriter
    );
  }
};


/**
 * optional messages.Patient patient = 1;
 * @return {?proto.heyrenee.v1.messages.Patient}
 */
proto.heyrenee.v1.UpdatePatientRequest.prototype.getPatient = function() {
  return /** @type{?proto.heyrenee.v1.messages.Patient} */ (
    jspb.Message.getWrapperField(this, heyrenee_v1_messages_patient_pb.Patient, 1));
};


/**
 * @param {?proto.heyrenee.v1.messages.Patient|undefined} value
 * @return {!proto.heyrenee.v1.UpdatePatientRequest} returns this
*/
proto.heyrenee.v1.UpdatePatientRequest.prototype.setPatient = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.UpdatePatientRequest} returns this
 */
proto.heyrenee.v1.UpdatePatientRequest.prototype.clearPatient = function() {
  return this.setPatient(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.UpdatePatientRequest.prototype.hasPatient = function() {
  return jspb.Message.getField(this, 1) != null;
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
proto.heyrenee.v1.CreateCaregiverRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.CreateCaregiverRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.CreateCaregiverRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.CreateCaregiverRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    caregiver: (f = msg.getCaregiver()) && heyrenee_v1_messages_caregiver_pb.Caregiver.toObject(includeInstance, f),
    password: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.heyrenee.v1.CreateCaregiverRequest}
 */
proto.heyrenee.v1.CreateCaregiverRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.CreateCaregiverRequest;
  return proto.heyrenee.v1.CreateCaregiverRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.CreateCaregiverRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.CreateCaregiverRequest}
 */
proto.heyrenee.v1.CreateCaregiverRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new heyrenee_v1_messages_caregiver_pb.Caregiver;
      reader.readMessage(value,heyrenee_v1_messages_caregiver_pb.Caregiver.deserializeBinaryFromReader);
      msg.setCaregiver(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPassword(value);
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
proto.heyrenee.v1.CreateCaregiverRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.CreateCaregiverRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.CreateCaregiverRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.CreateCaregiverRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCaregiver();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      heyrenee_v1_messages_caregiver_pb.Caregiver.serializeBinaryToWriter
    );
  }
  f = message.getPassword();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional messages.Caregiver caregiver = 1;
 * @return {?proto.heyrenee.v1.messages.Caregiver}
 */
proto.heyrenee.v1.CreateCaregiverRequest.prototype.getCaregiver = function() {
  return /** @type{?proto.heyrenee.v1.messages.Caregiver} */ (
    jspb.Message.getWrapperField(this, heyrenee_v1_messages_caregiver_pb.Caregiver, 1));
};


/**
 * @param {?proto.heyrenee.v1.messages.Caregiver|undefined} value
 * @return {!proto.heyrenee.v1.CreateCaregiverRequest} returns this
*/
proto.heyrenee.v1.CreateCaregiverRequest.prototype.setCaregiver = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.CreateCaregiverRequest} returns this
 */
proto.heyrenee.v1.CreateCaregiverRequest.prototype.clearCaregiver = function() {
  return this.setCaregiver(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.CreateCaregiverRequest.prototype.hasCaregiver = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional string password = 2;
 * @return {string}
 */
proto.heyrenee.v1.CreateCaregiverRequest.prototype.getPassword = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.CreateCaregiverRequest} returns this
 */
proto.heyrenee.v1.CreateCaregiverRequest.prototype.setPassword = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
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
proto.heyrenee.v1.GetCaregiverRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.GetCaregiverRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.GetCaregiverRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.GetCaregiverRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    caregiverId: jspb.Message.getFieldWithDefault(msg, 1, "")
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
 * @return {!proto.heyrenee.v1.GetCaregiverRequest}
 */
proto.heyrenee.v1.GetCaregiverRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.GetCaregiverRequest;
  return proto.heyrenee.v1.GetCaregiverRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.GetCaregiverRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.GetCaregiverRequest}
 */
proto.heyrenee.v1.GetCaregiverRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setCaregiverId(value);
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
proto.heyrenee.v1.GetCaregiverRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.GetCaregiverRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.GetCaregiverRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.GetCaregiverRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCaregiverId();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
};


/**
 * optional string caregiver_id = 1;
 * @return {string}
 */
proto.heyrenee.v1.GetCaregiverRequest.prototype.getCaregiverId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.heyrenee.v1.GetCaregiverRequest} returns this
 */
proto.heyrenee.v1.GetCaregiverRequest.prototype.setCaregiverId = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
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
proto.heyrenee.v1.UpdateCaregiverRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.heyrenee.v1.UpdateCaregiverRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.heyrenee.v1.UpdateCaregiverRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.UpdateCaregiverRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    caregiver: (f = msg.getCaregiver()) && heyrenee_v1_messages_caregiver_pb.Caregiver.toObject(includeInstance, f)
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
 * @return {!proto.heyrenee.v1.UpdateCaregiverRequest}
 */
proto.heyrenee.v1.UpdateCaregiverRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.heyrenee.v1.UpdateCaregiverRequest;
  return proto.heyrenee.v1.UpdateCaregiverRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.heyrenee.v1.UpdateCaregiverRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.heyrenee.v1.UpdateCaregiverRequest}
 */
proto.heyrenee.v1.UpdateCaregiverRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new heyrenee_v1_messages_caregiver_pb.Caregiver;
      reader.readMessage(value,heyrenee_v1_messages_caregiver_pb.Caregiver.deserializeBinaryFromReader);
      msg.setCaregiver(value);
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
proto.heyrenee.v1.UpdateCaregiverRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.heyrenee.v1.UpdateCaregiverRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.heyrenee.v1.UpdateCaregiverRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.heyrenee.v1.UpdateCaregiverRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getCaregiver();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      heyrenee_v1_messages_caregiver_pb.Caregiver.serializeBinaryToWriter
    );
  }
};


/**
 * optional messages.Caregiver caregiver = 1;
 * @return {?proto.heyrenee.v1.messages.Caregiver}
 */
proto.heyrenee.v1.UpdateCaregiverRequest.prototype.getCaregiver = function() {
  return /** @type{?proto.heyrenee.v1.messages.Caregiver} */ (
    jspb.Message.getWrapperField(this, heyrenee_v1_messages_caregiver_pb.Caregiver, 1));
};


/**
 * @param {?proto.heyrenee.v1.messages.Caregiver|undefined} value
 * @return {!proto.heyrenee.v1.UpdateCaregiverRequest} returns this
*/
proto.heyrenee.v1.UpdateCaregiverRequest.prototype.setCaregiver = function(value) {
  return jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.heyrenee.v1.UpdateCaregiverRequest} returns this
 */
proto.heyrenee.v1.UpdateCaregiverRequest.prototype.clearCaregiver = function() {
  return this.setCaregiver(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.heyrenee.v1.UpdateCaregiverRequest.prototype.hasCaregiver = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.heyrenee.v1);