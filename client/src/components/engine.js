import * as BABYLON from 'babylonjs';

export class B {
    engine = {};
    scene = {};
    camera = {}
    light = {}
    ground = {}
    tiles = {}
    pieces = {}
    pieceList = []
    canvas = {}
    startingPoint = {}
    currentMesh = {}


    static startup() {

        B.canvas = document.getElementById("renderCanvas");
        B.engine = new BABYLON.Engine(B.canvas, true, { preserveDrawingBuffer: true, stencil: true });
        B.createScene();

        B.engine.runRenderLoop(function () {
            if (B.scene) {
                B.scene.render();
            }
        });

        // Resize
        window.addEventListener("resize", function () {
            B.engine.resize();
        });

    }

    static createScene = () => {
        B.scene = new BABYLON.Scene(B.engine);
        B.camera = new BABYLON.ArcRotateCamera("Camera", 0, 0, 0, new BABYLON.Vector3(0, 0, 0), B.scene);

        B.camera.setPosition(new BABYLON.Vector3(0, 600, 50));

        B.camera.lowerBetaLimit = 0.1;
        B.camera.upperBetaLimit = (Math.PI / 2) * 0.99;
        B.camera.lowerRadiusLimit = 150;

        B.scene.clearColor = new BABYLON.Color3(0, 0, 0);

        // Light
        B.light = new BABYLON.PointLight("omni", new BABYLON.Vector3(80, 25, 80), B.scene);

        // Ground
        B.ground = BABYLON.Mesh.CreateGround("ground", 400, 300, 1, B.scene, false);
        const groundMaterial = new BABYLON.StandardMaterial("groundMaterial", B.scene);
        groundMaterial.specularColor = BABYLON.Color3.Black();
        B.ground.material = groundMaterial;

        const blackMat = new BABYLON.StandardMaterial("blackMat", B.scene);
        blackMat.diffuseColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        blackMat.specularColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        blackMat.emissiveColor = BABYLON.Color3.Black();
        const whiteMat = new BABYLON.StandardMaterial("whiteMat", B.scene);
        whiteMat.diffuseColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        whiteMat.specularColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        whiteMat.emissiveColor = BABYLON.Color3.Gray();

        let boards = 2;

        for (let board = 1; board <= boards; board++) {
            for (let x = 1; x <= 8; x++) {
                for (let y = 1; y <= 8; y++) {
                    const tile = BABYLON.Mesh.CreateBox("tile" + board + "." + x + "." + y, 20, B.scene);
                    tile.scaling = new BABYLON.Vector3(1, 0.1, 1);
                    if ((x + y) % 2) {
                        tile.material = blackMat;
                    } else {
                        tile.material = whiteMat;
                    }

                    tile.position.x = x * 20 - (board - 1) * 200 + 10;
                    tile.position.z = y * 20 - 90;
                    tile.position.y = 1;
                    tile.movable = false;

                }
            }
        }



        // Meshes
        B.redMat = new BABYLON.StandardMaterial("redMat", B.scene);
        B.redMat.diffuseColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.redMat.specularColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.redMat.emissiveColor = BABYLON.Color3.Red();

        let id = "piece1"
        if (!B.pieceExists(id)) {
            B.createPiece(id)
        }
        id = "piece2"
        if (!B.pieceExists(id)) {
            B.createPiece(id)
        }



        // Events
        B.canvas = B.engine.getRenderingCanvas();



        B.canvas.addEventListener("pointerdown", B.onPointerDown, false);
        B.canvas.addEventListener("pointerup", B.onPointerUp, false);
        B.canvas.addEventListener("pointermove", B.onPointerMove, false);

        B.scene.onDispose = () => {
            B.canvas.removeEventListener("pointerdown", B.onPointerDown);
            B.canvas.removeEventListener("pointerup", B.onPointerUp);
            B.canvas.removeEventListener("pointermove", B.onPointerMove);
        }

    };

    static getGroundPosition = () => {
        // Use a predicate to get position on the ground
        let pickinfo = B.scene.pick(B.scene.pointerX, B.scene.pointerY, function (mesh) { return mesh == B.ground; });
        if (pickinfo.hit) {
            return pickinfo.pickedPoint;
        }
        return null;
    }


    static onPointerDown = (evt) => {
        if (evt.button !== 0) {
            return;
        }

        // check if we are under a mesh
        const pickInfo = B.scene.pick(B.scene.pointerX, B.scene.pointerY, function (mesh) { return mesh !== B.ground; });
        if (pickInfo.hit) {
            if (pickInfo.pickedMesh.movable == false) {
                return; ///////////////////////////////////////////////////
            }
            B.currentMesh = pickInfo.pickedMesh;

            B.startingPoint = B.getGroundPosition(evt);

            if (B.startingPoint) { // we need to disconnect camera from canvas
                setTimeout(function () {
                    B.camera.detachControl(B.canvas);
                }, 0);
            }
        }
    }

    static onPointerUp = () => {

        // if (B.currentMesh.position.x)
        let x = B.currentMesh.position.x + 10;
        let newX = 20 * Math.round(x / 20);
        B.currentMesh.position.x = newX - 10;

        let z = B.currentMesh.position.z + 10;
        let newZ = 20 * Math.round(z / 20);
        B.currentMesh.position.z = newZ - 10;

        if (B && B.currentMesh) {
            console.log(B.currentMesh.metadata.id + ": " + B.currentMesh.position);
        }


        if (B.startingPoint) {
            B.camera.attachControl(B.canvas, true);
            B.startingPoint = null;
            return;
        }
    }

    static onPointerMove = (evt) => {
        if (!B.startingPoint) {
            return;
        }

        const current = B.getGroundPosition(evt);

        if (!current) {
            return;
        }

        const diff = current.subtract(B.startingPoint);
        B.currentMesh.position.addInPlace(diff);

        B.startingPoint = current;

    }

    static pieceExists(id) {
        if (B.pieces) {
            return B.pieces[id]
        }
        return false
    }

    static createPiece(id, color, identity) {
        let newPiece = BABYLON.Mesh.CreateBox(id, 20, B.scene);
        newPiece.material = B.redMat;
        newPiece.position.x += -100 + (8 * 20) + 110;
        newPiece.position.z += (8 * 20) + 10 - 100;
        newPiece.position.y = -9 + 20;
        newPiece.movable = true;
        newPiece.metadata = {}
        newPiece.metadata.id = id
        newPiece.metadata.color = color
        newPiece.metadata.identity = identity

        if (!B.pieces) {
            B.pieces = {}
        }
        B.pieces[id] = newPiece
        if (!B.pieceList) {
            B.pieceList = []
        }
        B.pieceList.push(newPiece)

    }




} 