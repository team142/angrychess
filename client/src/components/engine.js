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
    currfentMesh = {}

    NetworkManager


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

        B.blueMat = new BABYLON.StandardMaterial("blueMat", B.scene);
        B.blueMat.diffuseColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.blueMat.specularColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.blueMat.emissiveColor = BABYLON.Color3.Blue();

        B.redMat = new BABYLON.StandardMaterial("redMat", B.scene);
        B.redMat.diffuseColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.redMat.specularColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.redMat.emissiveColor = BABYLON.Color3.Red();



        B.blackMat = new BABYLON.StandardMaterial("blackMat", B.scene);
        B.blackMat.diffuseColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.blackMat.specularColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.blackMat.emissiveColor = BABYLON.Color3.Black();

        B.whiteMat = new BABYLON.StandardMaterial("whiteMat", B.scene);
        B.whiteMat.diffuseColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.whiteMat.specularColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.whiteMat.emissiveColor = BABYLON.Color3.White();

        B.grayMat = new BABYLON.StandardMaterial("grayMat", B.scene);
        B.grayMat.diffuseColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.grayMat.specularColor = new BABYLON.Color3(0.4, 0.4, 0.4);
        B.grayMat.emissiveColor = BABYLON.Color3.Gray();

        let boards = 2;

        for (let board = 1; board <= boards; board++) {
            for (let x = 1; x <= 8; x++) {
                for (let y = 1; y <= 8; y++) {
                    const tile = BABYLON.Mesh.CreateBox("tile" + board + "." + x + "." + y, 20, B.scene);
                    tile.scaling = new BABYLON.Vector3(1, 0.1, 1);
                    if (board % 2 == 0) {
                        if ((x + y) % 2) {
                            tile.material = B.grayMat;
                        } else {
                            tile.material = B.blackMat;
                        }
                    } else {
                        if ((x + y) % 2) {
                            tile.material = B.blackMat;
                        } else {
                            tile.material = B.grayMat;
                        }
                    }


                    tile.position.x = x * 20 - (board - 1) * 200 + 10;
                    tile.position.z = y * 20 - 90;
                    tile.position.y = 1;
                    tile.movable = false;

                }
            }
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



        if (B.startingPoint) {
            B.camera.attachControl(B.canvas, true);
            B.startingPoint = null;
            if (B && B.currentMesh) {
                // console.log(JSON.stringify(B.getPosition(B.currentMesh)));
                // console.log(B.currentMesh.metadata.id + ": " + B.currentMesh.position);
            }
            B.checkForMove(B.currentMesh)

            return;
        }



    }

    static checkForMove(mesh) {
        let p = B.getPosition(mesh)
        B.NetworkManager.sendMove(mesh.metadata.id, p.board, p.x, p.y, p.cache)
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

    static createPiece(piece) {

        let newPiece
        if (B.pieceList && (B.pieceExists.length > 0)) {
            newPiece = B.pieceList[0].clone(piece.id)
        } else {
            const cone = BABYLON.MeshBuilder.CreateCylinder("cone" + piece.id, { diameterTop: 0, height: 20, tessellation: 96, diameterBottom: 20 }, B.scene);
            const sphere = BABYLON.MeshBuilder.CreateSphere("sphere" + piece.id, { diameter: 10 }, B.scene);
            sphere.position.y = cone.position.y + 10
            newPiece = BABYLON.MeshBuilder.CreateBox(piece.id, { width: 20, height: 30, depth: 20 }, B.scene);
            newPiece.addChild(sphere)
            newPiece.addChild(cone)
            newPiece.visibility = false;
        }



        if (piece.color) {
            for (let a of newPiece.getChildMeshes()) {
                a.material = B.blueMat;
            }
        } else {
            for (let a of newPiece.getChildMeshes()) {
                a.material = B.redMat;
            }
        }



        newPiece.movable = true;
        newPiece.metadata = {}
        newPiece.metadata = piece

        if (!B.pieces) {
            B.pieces = {}
        }
        B.pieces[piece.id] = newPiece
        if (!B.pieceList) {
            B.pieceList = []
        }
        B.pieceList.push(newPiece)

    }

    static updatePiece(piece) {
        let newPiece = B.pieces[piece.id]
        newPiece.metadata = piece
        if (!piece.cache) {
            newPiece.position.x = 170 - (piece.x - 1) * 20 - 200 * (piece.board - 1);
            newPiece.position.z = 70 - (piece.y - 1) * 20;
            newPiece.position.y = -9 + 20;

        } else {
            B.updateCachedPiece(piece)
        }

    }

    static updateCachedPiece(piece) {
        let newPiece = B.pieces[piece.id]

        let n = B.getCachedRelativeNumber(piece.board, piece)
        let row = Math.floor((n - 1) / 10) + 1

        n = n - ((row - 1) * 10)
        let x = 190 - (n - 1) * 20;
        x = x - (piece.board - 1) * 200

        newPiece.position.x = x

        if ((piece.board % 2 == 1 && !piece.color) || (piece.board % 2 == 0 && piece.color)) {
            newPiece.position.z = -130 + 20 * (row - 1)
        } else {
            let z = 90 + 20 * (row - 1);
            newPiece.position.z = z;
        }

        newPiece.position.y = -9 + 20;
    }

    static getCachedRelativeNumber(board, piece) {
        let n = 0;
        for (const pieceB of B.pieceList) {
            if (pieceB.metadata.cache && pieceB.metadata.color == piece.color && pieceB.metadata.board == piece.board) {
                n++;
                if (pieceB.metadata.id == piece.id) {
                    return n
                }
            }
        }
        return n
    }


    static getPosition(mesh) {
        let result = {
            x: 0,
            y: 0,
            board: 0,
            cache: false
        }
        if (mesh && mesh.position) {
            if (mesh.position.x <= 170 && mesh.position.x >= 30 && mesh.position.z <= 70 && mesh.position.z >= -70) {
                result.cache = false
                result.board = 1
            } else if (mesh.position.x <= -30 && mesh.position.x >= -170 && mesh.position.z <= 70 && mesh.position.z >= -70) {
                result.cache = false
                result.board = 2
            } else {
                result.cache = true
                if (mesh.position.x >= 10) {
                    result.board = 1
                } else {
                    result.board = 2
                }

            }
            let x = mesh.position.x - 10
            x = 160 - x
            result.x = B.relativeTile(x)
            if (result.board == 2) {
                result.x -= 10;
            }
            let z = mesh.position.z - 10 + 100
            z = 160 - z
            result.y = B.relativeTile(z)
        }
        return result
    }

    static relativeTile(p) {
        if (p == 0) {
            return 1;
        }
        return Math.round(p / 20) + 1
    }

    static renderPlayer(player) {
        if (!player.pieces) {
            return
        }
        for (const piece of player.pieces) {
            B.createOrUpdatePiece(piece)
        }

    }

    static findMeshByIdentity(identity) {
        if (B.pieceList) {
            for (let piece of B.pieceList) {
                if (piece.metadata.identity == identity) {
                    return piece;
                }
            }
        }
        return null;
    }

    static createOrUpdatePiece(piece) {
        if (B.pieceExists(piece.id)) {
            B.updatePiece(piece)
        } else {
            B.createPiece(piece)
            B.updatePiece(piece)
        }
    }


} 